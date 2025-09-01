<script setup lang="ts">
import type { DataTableColumns } from 'naive-ui'
import type { TaskStatus } from '@/api/type'
import {
  CheckmarkCircleOutline,
  CloudDownloadOutline,
  DownloadOutline,
  ReloadOutline,
} from '@vicons/ionicons5'
import dayjs from 'dayjs'
import { NButton, useDialog, useMessage, useNotification } from 'naive-ui'

import { storeToRefs } from 'pinia'
import { onActivated, ref } from 'vue'
import { GetTaskProgress, RetryEncodeTask, RetryMergeTask } from '@/api'
import { useSettingStore } from '@/store/setting'
import { renderIcon, renderIconButton } from '@/util/render'

const props = defineProps({
  videoKey: {
    type: String,
    default: '',
  },
})

const { clipTimeout, queueName } = storeToRefs(useSettingStore())

const notification = useNotification()
const dialog = useDialog()
const message = useMessage()

interface TaskInfo {
  create_at: string
  encode_key: string
  encode_param: string
  encode_size: string
  encode_url: string
  key: string
  script: string
  size: string
  status: TaskStatus
  url: string
  percentage: number
}

interface TaskProgress {
  clip_key: string
  clip_url: string
  completed: boolean
  encode_key: string
  encode_url: string
  index: number
}

const columns: DataTableColumns<TaskProgress> = [
  {
    title: 'Index',
    key: 'index',
    fixed: 'left',
  },
  {
    title: 'Completed',
    key: 'completed',
    render: (row: TaskProgress) =>
      row.completed
        ? renderIcon(CheckmarkCircleOutline, {
            size: 20,
          })()
        : null,
  },
  {
    title: 'Retry Encode',
    key: 'retry',
    render: (row: TaskProgress) =>
      renderIconButton(ReloadOutline, () => handleRetryEncode(row.index)),
  },
  {
    title: 'Download Encode Clip',
    key: 'encode_url',
    render: (row: TaskProgress) =>
      row.encode_url
        ? renderIconButton(CloudDownloadOutline, () => window.open(row.encode_url, '_blank'))
        : null,
  },
  {
    title: 'Download Clip',
    key: 'clip_url',
    render: (row: TaskProgress) =>
      row.clip_url
        ? renderIconButton(DownloadOutline, () => window.open(row.clip_url, '_blank'))
        : null,
  },
]

const taskInfo = ref<TaskInfo | null>(null)
const taskProgress = ref<TaskProgress[]>([])
onActivated(() => {
  fetchTaskProgress()
})
function fetchTaskProgress(): void {
  console.log('Fetch task progress...')
  GetTaskProgress({
    video_key: props.videoKey,
  })
    .then((res) => {
      if (res.success) {
        if (!res.data) {
          console.error('Task progress is null')
          return
        }

        let total: number = 0
        let done: number = 0

        const tempProgress: TaskProgress[] = []
        res.data?.progress.forEach((task) => {
          total++
          if (task.completed) {
            done++
          }

          tempProgress.push({
            clip_key: task.clip_key,
            clip_url: task.clip_url,
            completed: task.completed,
            encode_key: task.encode_key,
            encode_url: task.encode_url,
            index: task.index,
          })
        })

        const tempInfo: TaskInfo = {
          create_at: dayjs.unix(res.data.create_at).format('YYYY-MM-DD HH:mm:ss'),
          encode_key: res.data.encode_key,
          encode_param: res.data.encode_param,
          encode_size: res.data.encode_size,
          encode_url: res.data.encode_url,
          key: res.data.key,
          script: res.data.script,
          size: res.data.size,
          status: res.data.status,
          url: res.data.url,
          percentage: Number(((done / total) * 100).toFixed(2)),
        }

        taskProgress.value = tempProgress
        taskInfo.value = tempInfo
      }
      else {
        console.error(res)
        notification.error({
          content: 'Fetch task progress failed',
          meta: String(res) || 'Unknown error',
        })
      }
    })
    .catch((err) => {
      console.error(err)
      notification.error({
        content: 'Fetch task progress failed',
        meta: String(err) || 'Unknown error',
      })
    })
}

function downloadVideo(encode: boolean): void {
  const url = encode ? taskInfo.value?.encode_url : taskInfo.value?.url
  window.open(url, '_blank')
}

function handleRetryMerge(): void {
  dialog.warning({
    title: 'Retry Merge Task?',
    positiveText: 'MERGE',
    negativeText: 'NO',
    maskClosable: false,
    onMaskClick: () => {
      message.warning('Cannot close')
    },
    onPositiveClick: () => {
      RetryMergeTask({
        video_key: String(taskInfo.value?.key),
      })
        .then((res) => {
          if (res.success) {
            notification.success({
              content: 'Merging...',
              meta: `Task: ${String(taskInfo.value?.key)}`,
              duration: 2500,
              keepAliveOnHover: true,
            })
          }
          else {
            notification.error({
              content: 'Start merge failed',
              meta: res.error?.message || 'Unknown error',
            })
          }
        })
        .catch((err) => {
          console.error(err)
          notification.error({
            content: 'Start merge failed',
            meta: String(err) || 'Unknown error',
          })
        })
        .finally(() => {
          fetchTaskProgress()
        })
    },
  })
}

function handleRetryEncode(index: number): void {
  dialog.warning({
    title: 'Retry Encode Task?',
    content: 'Are you sure to retry encoding this clip?',
    positiveText: 'ENCODE',
    negativeText: 'NO',
    maskClosable: false,
    onMaskClick: () => {
      message.warning('Cannot close')
    },
    onPositiveClick: () => {
      RetryEncodeTask({
        video_key: String(taskInfo.value?.key),
        index,
        timeout: clipTimeout.value,
        queue: queueName.value,
      })
        .then((res) => {
          if (res.success) {
            notification.success({
              content: 'Retry encoding...',
              meta: `Clip: ${String(index)}`,
              duration: 2500,
              keepAliveOnHover: true,
            })
          }
          else {
            notification.error({
              content: 'Start re-encode failed',
              meta: res.error?.message || 'Unknown error',
            })
          }
        })
        .catch((err) => {
          console.error(err)
          notification.error({
            content: 'Start re-encode failed',
            meta: String(err) || 'Unknown error',
          })
        })
        .finally(() => {
          fetchTaskProgress()
        })
    },
  })
}
</script>

<template>
  <NSpace vertical>
    <NCard :title="taskInfo?.key">
      <NSpace>
        <NSpace>
          <div v-if="taskInfo?.status !== 'completed'">
            <NProgress type="circle" :percentage="taskInfo?.percentage" />
          </div>
          <div v-else>
            <NProgress type="circle" status="success" :percentage="100" />
          </div>
        </NSpace>
        <NSpace vertical class="p-5">
          <NButton type="warning" @click="downloadVideo(false)">
            Download Origin ({{ taskInfo?.size }})
          </NButton>
          <div v-if="taskInfo?.status === 'completed'">
            <NButton type="success" @click="downloadVideo(true)">
              Download Encode ({{ taskInfo?.encode_size }})
            </NButton>
          </div>
        </NSpace>
      </NSpace>
    </NCard>
    <NCard title="Script">
      <NCode :code="taskInfo?.script" language="python" show-line-numbers />
    </NCard>
    <NCard title="Encode Param">
      <NCode :code="taskInfo?.encode_param" word-wrap />
    </NCard>
    <NCard>
      <NSpace vertical>
        <NSpace justify="space-between">
          <NGradientText size="18" gradient="linear-gradient(90deg, red 0%, green 50%, blue 100%)">
            Clips
          </NGradientText>
          <NButton type="warning" @click="handleRetryMerge">
            Retry Merge
          </NButton>
        </NSpace>
        <NDataTable
          :columns="columns"
          :data="taskProgress"
          :row-key="(row: TaskProgress) => row.index"
          max-height="100vh"
          virtual-scroll
          striped
        />
      </NSpace>
    </NCard>
  </NSpace>
</template>

<style scoped></style>

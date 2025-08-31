<script setup lang="ts">
import { DownloadOutline, PlayCircleOutline } from '@vicons/ionicons5'
import dayjs from 'dayjs'
import type { DataTableColumns, DataTableRowKey } from 'naive-ui'
import { NButton, useDialog, useMessage, useNotification } from 'naive-ui'
import { storeToRefs } from 'pinia'

import { ClearTask, GetTaskList, StartTask } from '@/api'
import { useSettingStore } from '@/store/setting'
import { renderIconButton } from '@/util/render'

const { script, encodeParam, sliceMode, clipTimeout, queueName } = storeToRefs(useSettingStore())

const notification = useNotification()
const dialog = useDialog()
const message = useMessage()

interface pendingTask {
  key: string
  create_at: string
  url: string
}

const columns: DataTableColumns<pendingTask> = [
  {
    type: 'selection',
    fixed: 'left',
  },
  {
    title: 'Video',
    key: 'key',
  },
  {
    title: 'Run',
    key: 'run',
    render: (row: pendingTask) => renderIconButton(PlayCircleOutline, () => submitTasks([row.key])),
  },
  {
    title: 'Date',
    key: 'create_at',
  },
  {
    title: 'Download Origin',
    key: 'download',
    render: (row: pendingTask) =>
      row.url ? renderIconButton(DownloadOutline, () => window.open(row.url, '_blank')) : null,
  },
]

const tasks = ref<pendingTask[]>([])
onActivated(() => {
  fetchPendingTasks()
})
function fetchPendingTasks(): void {
  console.log('Fetch pending tasks...')
  GetTaskList({
    pending: true,
    running: false,
    completed: false,
  })
    .then((res) => {
      if (res.success) {
        const temp: pendingTask[] = []
        res.data?.forEach((task) => {
          temp.push({
            key: task.key,
            create_at: dayjs.unix(task.create_at).format('YYYY-MM-DD HH:mm:ss'),
            url: task.url,
          })
        })
        tasks.value = temp
      } else {
        console.error(res)
        notification['error']({
          content: 'Fetch task list failed',
          meta: String(res) || 'Unknown error',
        })
      }
    })
    .catch((err) => {
      console.error(err)
      notification['error']({
        content: 'Fetch task list failed',
        meta: String(err) || 'Unknown error',
      })
    })
}

const checkedRowKeys = ref<DataTableRowKey[]>([])
function updateCheckedRowKeys(keys: DataTableRowKey[]): void {
  checkedRowKeys.value = keys
}

function submitTasks(taskKeys: DataTableRowKey[]): void {
  if (taskKeys.length === 0) {
    message.warning('Please select at least one task')
    return
  }

  dialog.info({
    title: taskKeys.length > 1 ? 'Run Selected Tasks?' : 'Run Task?',
    content: 'Are you sure to start with the latest script and encode param?',
    positiveText: 'RUN',
    negativeText: 'NO',
    maskClosable: false,
    onMaskClick: () => {
      message.warning('Cannot close')
    },
    onPositiveClick: () => {
      taskKeys.forEach((key) => {
        console.log(key.toString())
        StartTask({
          encode_param: encodeParam.value,
          script: script.value,
          video_key: key.toString(),
          slice: sliceMode.value,
          timeout: clipTimeout.value,
          queue: queueName.value,
        })
          .then((res) => {
            if (res.success) {
              notification['success']({
                content: 'Task started successfully',
                meta: 'Task: ' + key,
                duration: 2500,
                keepAliveOnHover: true,
              })
            } else {
              notification['error']({
                content: 'Task start failed',
                meta: res.error?.message || 'Unknown error',
              })
            }

            updateCheckedRowKeys(checkedRowKeys.value.filter((k) => k !== key))
          })
          .catch((err) => {
            console.error(err)
            notification['error']({
              content: 'Task start failed',
              meta: String(err) || 'Unknown error',
            })
          })
          .finally(() => {
            fetchPendingTasks()
          })
      })
    },
  })
}

function deleteTasks(taskKeys: DataTableRowKey[]): void {
  if (taskKeys.length === 0) {
    message.warning('Please select at least one task')
    return
  }

  dialog.warning({
    title: taskKeys.length > 1 ? 'Delete Selected Tasks?' : 'Delete Task?',
    positiveText: 'DELETE',
    negativeText: 'NO',
    maskClosable: false,
    onMaskClick: () => {
      message.warning('Cannot close')
    },
    onPositiveClick: () => {
      taskKeys.forEach((key) => {
        ClearTask({
          video_key: key.toString(),
        })
          .then((res) => {
            if (res.success) {
              notification['success']({
                content: 'Task deleted successfully',
                meta: 'Task: ' + key,
                duration: 2500,
                keepAliveOnHover: true,
              })
            } else {
              notification['error']({
                content: 'Delete task failed',
                meta: res.error?.message || 'Unknown error',
              })
            }

            updateCheckedRowKeys(checkedRowKeys.value.filter((k) => k !== key))
          })
          .catch((err) => {
            console.error(err)
            notification['error']({
              content: 'Delete task failed',
              meta: String(err) || 'Unknown error',
            })
          })
          .finally(() => {
            fetchPendingTasks()
          })
      })
    },
  })
}
</script>

<template>
  <NCard>
    <NSpace vertical>
      <NSpace justify="space-between">
        <NGradientText size="18" type="warning"> Pending </NGradientText>
        <NSpace>
          <NButton type="error" @click="deleteTasks(checkedRowKeys)"> Delete </NButton>
          <NButton type="primary" @click="submitTasks(checkedRowKeys)"> RUN </NButton>
        </NSpace>
      </NSpace>
      <NDataTable
        :columns="columns"
        :data="tasks"
        :row-key="(row: pendingTask) => row.key"
        max-height="70vh"
        virtual-scroll
        striped
        :checked-row-keys="checkedRowKeys"
        @update:checked-row-keys="updateCheckedRowKeys"
      />
    </NSpace>
  </NCard>
</template>

<style scoped></style>

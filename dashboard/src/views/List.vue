<script setup lang="ts">
import { CloudDownloadOutline, SearchOutline } from '@vicons/ionicons5'
import dayjs from 'dayjs'
import type { DataTableColumns, DataTableRowKey } from 'naive-ui'
import { NButton, useDialog, useMessage, useNotification } from 'naive-ui'
import { storeToRefs } from 'pinia'

import { ClearTask, GetTaskList } from '@/api'
import type { TaskStatus } from '@/api/type'
import { useSettingStore } from '@/store/setting'
import { renderIconButton, renderStatusButton } from '@/util/render'

const { checkedPendingBox, checkedRunningBox, checkedCompletedBox } = storeToRefs(useSettingStore())

const notification = useNotification()
const dialog = useDialog()
const message = useMessage()

interface Task {
  create_at: string
  encode_key: string
  encode_param: string
  encode_url: string
  key: string
  script: string
  status: TaskStatus
  url: string
}

const columns: DataTableColumns<Task> = [
  {
    type: 'selection',
    fixed: 'left',
  },
  {
    title: 'Video',
    key: 'key',
    fixed: 'left',
    filter: (value: string | number, row: Task) =>
      row.key.toLowerCase().includes(String(value).toLowerCase()),
  },
  {
    title: 'Status',
    key: 'status',
    render: (row: Task) => renderStatusButton(row.status, row.key),
  },
  {
    title: 'Date',
    key: 'create_at',
  },
  {
    title: 'Download Encode',
    key: 'download',
    render: (row: Task) =>
      row.encode_url
        ? renderIconButton(CloudDownloadOutline, () => window.open(row.encode_url, '_blank'))
        : null,
  },
]

const tasks = ref<Task[]>([])
onActivated(() => {
  fetchTasks()
})
function fetchTasks(): void {
  console.log('Fetch pending tasks...')
  GetTaskList({
    pending: checkedPendingBox.value,
    running: checkedRunningBox.value,
    completed: checkedCompletedBox.value,
  })
    .then((res) => {
      if (res.success) {
        const temp: Task[] = []
        res.data?.forEach((task) => {
          temp.push({
            create_at: dayjs.unix(task.create_at).format('YYYY-MM-DD HH:mm:ss'),
            encode_key: task.encode_key,
            encode_param: task.encode_param,
            encode_url: task.encode_url,
            key: task.key,
            script: task.script,
            status: task.status,
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

function deleteTasks(taskKeys: DataTableRowKey[]): void {
  if (taskKeys.length === 0) {
    message.warning('Please select at least one task')
    return
  }

  dialog.warning({
    title: 'Delete Selected Task?',
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
            fetchTasks()
          })
      })
    },
  })
}

function downloadTasks(taskKeys: DataTableRowKey[]): void {
  if (taskKeys.length === 0) {
    message.warning('Please select at least one task')
    return
  }

  let downloadList: string[] = []
  taskKeys.forEach((key) => {
    const task = tasks.value.find((t) => t.key === key)
    if (task?.encode_url) {
      downloadList.push(task.encode_url)
    }
  })

  if (downloadList.length === 0) {
    message.warning('No download link found')
    return
  }

  downloadList.forEach((url) => {
    window.open(url, '_blank')
  })
}

const filter = ref('')
const tableRef: any = ref(null)
watch(filter, () => {
  console.log('Filter Update:', filter.value)
  tableRef.value.filter({
    key: filter.value,
  })
})
</script>

<template>
  <NCard>
    <NSpace vertical>
      <NSpace justify="space-between">
        <NSpace item-style="display: flex;" align="center">
          <NCheckbox v-model:checked="checkedPendingBox" @update-checked="fetchTasks">
            <NGradientText type="warning"> Pending</NGradientText>
          </NCheckbox>
          <NCheckbox v-model:checked="checkedRunningBox" @update-checked="fetchTasks">
            <NGradientText type="info"> Running</NGradientText>
          </NCheckbox>
          <NCheckbox v-model:checked="checkedCompletedBox" @update-checked="fetchTasks">
            <NGradientText type="success"> Completed</NGradientText>
          </NCheckbox>
          <NInput v-model:value="filter" placeholder="Search">
            <template #prefix>
              <NIcon :component="SearchOutline" />
            </template>
          </NInput>
        </NSpace>
        <NSpace>
          <NButton type="error" @click="deleteTasks(checkedRowKeys)"> Delete </NButton>
          <NButton type="info" @click="downloadTasks(checkedRowKeys)"> Download </NButton>
        </NSpace>
      </NSpace>
      <NDataTable
        ref="tableRef"
        :columns="columns"
        :data="tasks"
        :row-key="(row: Task) => row.key"
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

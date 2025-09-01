<script setup lang="ts">
import { useNotification } from 'naive-ui'
import { storeToRefs } from 'pinia'

import { onMounted } from 'vue'
import { Ping } from '@/api'
import { getGitHubTemplates } from '@/api/github'
import router from '@/router'
import { useSettingStore } from '@/store/setting'
import List from '@/views/List.vue'

const { templateRepo, vsScriptTemplates, ffmpegParamTemplates, githubToken }
  = storeToRefs(useSettingStore())

const notification = useNotification()

onMounted(() => {
  Ping()
    .then((res) => {
      if (!res.success) {
        router.push('/setting')
        notification.error({
          content: 'Server is not available',
          meta: res.error?.message || 'Unknown error',
          duration: 2500,
          keepAliveOnHover: true,
        })
      }
    })
    .catch((error) => {
      router.push('/setting')
      notification.error({
        content: 'Server is not available',
        meta: String(error) || 'Unknown error',
      })
    })

  getGitHubTemplates(templateRepo.value, githubToken.value, 'py')
    .then((res) => {
      console.log(res)
      vsScriptTemplates.value = res
    })
    .catch((error) => {
      console.log(error)
    })

  getGitHubTemplates(templateRepo.value, githubToken.value, 'txt')
    .then((res) => {
      console.log(res)
      ffmpegParamTemplates.value = res
    })
    .catch((error) => {
      console.log(error)
    })
})
</script>

<template>
  <List />
</template>

<style scoped></style>

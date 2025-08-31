<script setup lang="ts">
import hljs from 'highlight.js/lib/core'
import python from 'highlight.js/lib/languages/python'
import { darkTheme, lightTheme, useOsTheme } from 'naive-ui'
import { storeToRefs } from 'pinia'

import Layout from '@/layout/Layout.vue'
import { useSettingStore } from '@/store/setting'
import { themeOverrides } from '@/theme'

hljs.registerLanguage('python', python)

const { systemDarkMode, darkMode } = storeToRefs(useSettingStore())

const osThemeRef = useOsTheme()
watchEffect(() => {
  console.log('osTheme change!', osThemeRef.value)
  systemDarkMode.value = osThemeRef.value === 'dark'
})

const theme = computed(() => {
  if (darkMode.value === 'system') {
    return systemDarkMode.value ? darkTheme : lightTheme
  }
  return darkMode.value === 'dark' ? darkTheme : lightTheme
})
</script>

<template>
  <NConfigProvider :theme="theme" :theme-overrides="themeOverrides" :hljs="hljs">
    <NLoadingBarProvider>
      <NDialogProvider>
        <NMessageProvider>
          <NNotificationProvider>
            <Layout />
          </NNotificationProvider>
        </NMessageProvider>
      </NDialogProvider>
    </NLoadingBarProvider>
  </NConfigProvider>
</template>

<script setup lang="ts">
import { VueMonacoEditor } from '@guolao/vue-monaco-editor'
import * as monacoEditor from 'monaco-editor/esm/vs/editor/editor.api'
import type { SelectOption } from 'naive-ui'
import { useNotification } from 'naive-ui'
import { storeToRefs } from 'pinia'
import { shallowRef } from 'vue'

import { getGitHubTemplateContent } from '@/api/github'
import { useSettingStore } from '@/store/setting'

const { darkMode, systemDarkMode, script, encodeParam, vsScriptTemplates, ffmpegParamTemplates } =
  storeToRefs(useSettingStore())

const notification = useNotification()

const MONACO_EDITOR_OPTIONS: monacoEditor.editor.IStandaloneEditorConstructionOptions = {
  acceptSuggestionOnCommitCharacter: true,
  acceptSuggestionOnEnter: 'smart',
  automaticLayout: true,
  formatOnType: true,
  formatOnPaste: true,
}

const theme = computed(() => {
  if (darkMode.value === 'system') {
    return systemDarkMode.value ? 'vs-dark' : 'vs'
  }
  return darkMode.value === 'dark' ? 'vs-dark' : 'vs'
})

const editor = shallowRef()
function handleMount(editorInstance: any): any {
  editor.value = editorInstance
}

function handleUpdateVSScriptTemplate(value: string, option: SelectOption): void {
  console.log('fetching template content: ' + value)
  getGitHubTemplateContent(option)
    .then((res) => {
      script.value = res
    })
    .catch((err) => {
      console.error(err)
      notification['error']({
        title: 'Failed to get vs script template content',
        content: String(err),
      })
    })
}

function handleUpdateFFmpegParamTemplate(value: string, option: SelectOption): void {
  console.log('fetching template content: ' + value)
  getGitHubTemplateContent(option)
    .then((res) => {
      // if the content ends with line break or space, remove them
      encodeParam.value = res.trimEnd()
    })
    .catch((err) => {
      console.error(err)
      notification['error']({
        title: 'Failed to get ffmpeg encode param template content',
        content: String(err),
      })
    })
}

function handleEncodeParamChange(value: string): void {
  console.log('encode param changed: ' + value)
  // 如果输入的字符串里面包含了换行符，提醒用户
  if (value.includes('\n') || value.includes('\r')) {
    notification['error']({
      title: 'Encode Param cannot contain line break!!!',
      content: 'Please remove line breaks',
    })
  }
}
</script>

<template>
  <NSpace vertical>
    <NSpace justify="space-between">
      <NGradientText size="18" type="primary"> Code </NGradientText>
      <NSelect
        :options="vsScriptTemplates"
        style="width: 50vh"
        @update:value="handleUpdateVSScriptTemplate"
      />
    </NSpace>
    <NCard hoverable size="small" style="width: 100%; height: 65vh">
      <VueMonacoEditor
        v-model:value="script"
        language="python"
        :theme="theme"
        :options="MONACO_EDITOR_OPTIONS"
        class="h-screen w-screen"
        @mount="handleMount"
      />
    </NCard>
    <NSpace justify="space-between">
      <NGradientText size="18" type="primary"> Encode Param </NGradientText>
      <NSelect
        :options="ffmpegParamTemplates"
        style="width: 50vh"
        @update:value="handleUpdateFFmpegParamTemplate"
      />
    </NSpace>
    <NCard hoverable size="small" style="width: 100%">
      <NInput
        v-model:value="encodeParam"
        type="textarea"
        placeholder="Encode Param"
        @change="handleEncodeParamChange"
      />
    </NCard>
  </NSpace>
</template>

<style scoped></style>

<script setup lang="ts">
import { DashboardOutlined } from '@vicons/antd'
import { CloudUploadOutline, Code, PlayCircleOutline } from '@vicons/ionicons5'
import type { MenuOption } from 'naive-ui'
import { h, ref } from 'vue'
import { RouterLink } from 'vue-router'

import logo from '@/assets/logo.svg'
import { renderIcon } from '@/util/render'

const props = defineProps({
  collapsed: {
    type: Boolean,
    default: false,
  },
})
const activeKey = ref<string | null>(null)

const menuOptions: MenuOption[] = [
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/',
          },
        },
        { default: () => 'Dashboard' },
      ),
    key: 'Dashboard',
    icon: renderIcon(DashboardOutlined),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/start',
          },
        },
        { default: () => 'Start' },
      ),
    key: 'Start',
    icon: renderIcon(PlayCircleOutline),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/upload',
          },
        },
        { default: () => 'Upload' },
      ),
    key: 'Upload',
    icon: renderIcon(CloudUploadOutline),
  },
  {
    label: () =>
      h(
        RouterLink,
        {
          to: {
            path: '/code',
          },
        },
        { default: () => 'Code' },
      ),
    key: 'Code',
    icon: renderIcon(Code),
  },
]
</script>

<template>
  <div class="h-screen">
    <div class="pt-[2vh] h-[6vh]">
      <div v-if="!props.collapsed" class="ml-14">
        <NGradientText size="18"> FinalRip </NGradientText>
      </div>
      <div v-else class="ml-5 mt-[-5px]">
        <NImage :src="logo" width="28" preview-disabled />
      </div>
    </div>
    <NMenu
      v-model:value="activeKey"
      :collapsed="props.collapsed"
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :options="menuOptions"
    />
  </div>
</template>

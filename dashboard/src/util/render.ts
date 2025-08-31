import { NButton, NIcon } from 'naive-ui'
import type { Component, VNodeChild } from 'vue'
import { h } from 'vue'

import type { TaskStatus } from '@/api/type'
import router from '@/router'

export function renderIcon(icon: Component, props?: any): () => VNodeChild {
  return () => h(NIcon, props, { default: () => h(icon) })
}

export function renderIconButton(icon: Component, onClick: () => void): VNodeChild {
  return h(
    NButton,
    {
      text: true,
      onClick: onClick,
    },
    {
      default: renderIcon(icon, {
        size: 20,
      }),
    },
  )
}

export function renderStatusButton(status: TaskStatus, videoKey: string): VNodeChild {
  let type: 'default' | 'error' | 'primary' | 'info' | 'success' | 'warning' | undefined
  if (status === 'pending') {
    type = 'warning'
  } else if (status === 'running') {
    type = 'info'
  } else if (status === 'completed') {
    type = 'success'
  } else {
    type = 'error'
  }

  const handleClick = (): void => {
    if (status === 'pending') {
      return
    }
    router.push('/task/' + videoKey)
  }

  return h(
    NButton,
    {
      secondary: true,
      bordered: true,
      type: type,
      onClick: handleClick,
    },
    {
      default: () => status,
    },
  )
}

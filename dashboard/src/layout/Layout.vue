<script setup lang="ts">
import { storeToRefs } from 'pinia'

import { useSettingStore } from '@/store/setting'

const { collapsed } = storeToRefs(useSettingStore())
</script>

<template>
  <NLayout has-sider position="absolute" :native-scrollbar="false">
    <NLayoutSider
      collapse-mode="width"
      :collapsed-width="64"
      :width="200"
      :collapsed="collapsed"
      show-trigger
      :native-scrollbar="false"
      bordered
      @collapse="collapsed = true"
      @expand="collapsed = false"
    >
      <Menu v-model:collapsed="collapsed" />
    </NLayoutSider>
    <NLayout :native-scrollbar="false">
      <NLayoutHeader bordered class="p-[2vh] h-[6vh]">
        <Header />
      </NLayoutHeader>
      <NLayoutContent>
        <NLayout :native-scrollbar="false" class="p-[3vh] h-[94vh]">
          <RouterView v-slot="{ Component }">
            <KeepAlive>
              <component :is="Component" />
            </KeepAlive>
          </RouterView>
        </NLayout>
      </NLayoutContent>
    </NLayout>
  </NLayout>
</template>

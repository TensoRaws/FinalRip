import { loader } from '@guolao/vue-monaco-editor'
import * as monaco from 'monaco-editor'
import editorWorker from 'monaco-editor/esm/vs/editor/editor.worker?worker'
import cssWorker from 'monaco-editor/esm/vs/language/css/css.worker?worker'
import htmlWorker from 'monaco-editor/esm/vs/language/html/html.worker?worker'
import jsonWorker from 'monaco-editor/esm/vs/language/json/json.worker?worker'
import tsWorker from 'monaco-editor/esm/vs/language/typescript/ts.worker?worker'
import {
  // create naive ui
  create,
  // component
  NButton,
  NCard,
  NCheckbox,
  NCode,
  NDataTable,
  NDivider,
  NDrawer,
  NDrawerContent,
  NFlex,
  NGradientText,
  NIcon,
  NImage,
  NInput,
  NInputNumber,
  NLayout,
  NLayoutContent,
  NLayoutFooter,
  NLayoutHeader,
  NLayoutSider,
  NLog,
  NMenu,
  NP,
  NPopover,
  NProgress,
  NSelect,
  NSpace,
  NSwitch,
  NText,
  NUpload,
  NUploadDragger,
} from 'naive-ui'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import '@/styles/tailwind.css'

// eslint-disable-next-line no-restricted-globals
self.MonacoEnvironment = {
  getWorker(_, label): any {
    if (label === 'json') {
      // eslint-disable-next-line new-cap
      return new jsonWorker()
    }
    if (label === 'css' || label === 'scss' || label === 'less') {
      // eslint-disable-next-line new-cap
      return new cssWorker()
    }
    if (label === 'html' || label === 'handlebars' || label === 'razor') {
      // eslint-disable-next-line new-cap
      return new htmlWorker()
    }
    if (label === 'typescript' || label === 'javascript') {
      // eslint-disable-next-line new-cap
      return new tsWorker()
    }
    // eslint-disable-next-line new-cap
    return new editorWorker()
  },
}

// 配置从 `node_modules` 中加载 monaco-editor
loader.config({ monaco })

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

const meta = document.createElement('meta')
meta.name = 'naive-ui-style'
document.head.appendChild(meta)

const naive = create({
  components: [
    NButton,
    NCard,
    NCheckbox,
    NCode,
    NDataTable,
    NDivider,
    NDrawer,
    NDrawerContent,
    NFlex,
    NGradientText,
    NIcon,
    NImage,
    NInput,
    NInputNumber,
    NLayout,
    NLayoutContent,
    NLayoutFooter,
    NLayoutHeader,
    NLayoutSider,
    NLog,
    NMenu,
    NP,
    NPopover,
    NProgress,
    NSelect,
    NSpace,
    NSwitch,
    NText,
    NUpload,
    NUploadDragger,
  ],
})

createApp(App).use(naive).use(pinia).use(router).mount('#app')

import { createApp } from 'vue'
import 'vue-global-api'
import App from './App.vue'
import { setupRouter } from 'Plugins/router'
import { setupStore } from 'Store/index'
// 样式引入
import 'ant-design-vue/dist/reset.css'
import 'uno.css'

const app = createApp(App)
async function setupApp() {
  // 挂载pinia状态管理
  setupStore(app)

  // 挂载路由
  await setupRouter(app)

  app.mount('#app')
}

setupApp()

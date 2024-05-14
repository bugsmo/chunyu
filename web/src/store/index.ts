import { createPinia } from 'pinia'
import type { App } from 'vue'

// 创建store实例
const store = createPinia()

// 挂载到app上
export function setupStore(app: App<Element>) {
  app.use(store)
}

export { store }

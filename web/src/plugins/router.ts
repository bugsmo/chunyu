import { App } from 'vue'
import { routerBeforeEachFunc, routerAfterEachFunc } from 'Config/interceptors/router' // 在vite.config.ts和tsconfig.json配置好路径别名Config
import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'

type RouterModule = {
  default: RouteRecordRaw[]
}

// 导入所有路由文件
let Routes: Array<RouteRecordRaw> = []

const modules = import.meta.glob('Routes/*.ts', { eager: true }) // 在vite.config.ts和tsconfig.json配置好路径别名Routes。eager: true表示同步导入

for (const path in modules) {
  const module: RouterModule = modules[path] as RouterModule
  Routes = Routes.concat(module.default)
}

// 创建路由实例
const router = createRouter({
  history: createWebHistory(),
  routes: Routes,
})

export async function setupRouter(app: App) {
  // 创建路由守卫
  router.beforeEach(routerBeforeEachFunc)
  router.afterEach(routerAfterEachFunc)

  app.use(router)

  // 路由准备就绪后挂载APP实例
  await router.isReady()
}

export default router

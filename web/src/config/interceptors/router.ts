import { RouteLocationNormalized, NavigationGuardNext } from 'vue-router'

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export async function routerBeforeEachFunc(to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) {
  // console.log(to, from)
  // 路由进入前的操作
  next()
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
export function routerAfterEachFunc(to: RouteLocationNormalized, from: RouteLocationNormalized) {
  // 路由进入后的操作
  // console.log(to, from)
}

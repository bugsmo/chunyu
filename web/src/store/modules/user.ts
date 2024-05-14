import { defineStore } from 'pinia'
import { $authLogin } from 'API/user'

// 第一个参数是该 store 的唯一 id
const userStore = defineStore('user', {
  state: () => {
    return {
      isLogined: 0, // 0 未知  1  登录  2 未登录
      username: 'test',
      email: 'test@test.com',
    }
  },
  actions: {
    authLogin() {
      return $authLogin().then(
        (res) => {
          this.isLogined = 1
          Object.assign(this.username, res.data.username)
          return res
        },
        (rej) => {
          this.isLogined = 2
          console.log('未登录')
          return rej
        }
      )
    },
  },
  // other options...
})

export default userStore

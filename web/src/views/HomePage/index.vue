<template>
  <a-layout class="layout">
    <a-layout-header>
      <div class="logo" />
      <a-menu v-model:selectedKeys="selectedKeys" theme="dark" mode="horizontal" :style="{ lineHeight: '64px' }">
        <a-menu-item key="1">nav 1</a-menu-item>
        <a-menu-item key="2">nav 2</a-menu-item>
        <a-menu-item key="3">nav 3</a-menu-item>
      </a-menu>
    </a-layout-header>
    <a-layout-content style="padding: 0 50px">
      <a-breadcrumb style="margin: 16px 0">
        <a-breadcrumb-item>Home</a-breadcrumb-item>
        <a-breadcrumb-item>List</a-breadcrumb-item>
        <a-breadcrumb-item>App</a-breadcrumb-item>
      </a-breadcrumb>
      <div :style="{ background: '#fff', minHeight: '280px' }" class="pa-100px">
        <div class="i-openmoji:home-button" />
        <i-openmoji:home-button />
        <div class="my-card">
          <div class="card-item">Card 1</div>
          <div class="card-item">Card 2</div>
          <div class="card-item">Card 3</div>
        </div>
      </div>
    </a-layout-content>
    <a-layout-footer style="text-align: center"> Ant Design ©2018 Created by Ant UED </a-layout-footer>
  </a-layout>
</template>
<script lang="ts" setup>
import { ref } from 'vue'
// 测试接口请求
import { $authLogin } from 'API/user' // 这里记得配置路径映射别名
// // 测试Store
// import useUserStore from 'Store/modules/user'

// const userStore = useUserStore()

// console.log(userStore.username)

// userStore.authLogin().then((res) => {
//   console.log(res)
// })

$authLogin({
  username: 'test',
  password: 'test',
}).then((res) => {
  console.log(res)
})

const selectedKeys = ref<string[]>(['2'])
</script>
<style scoped>
.site-layout-content {
  padding: 24px;
  min-height: 280px;
  background: #fff;
}

#components-layout-demo-top .logo {
  float: left;
  margin: 16px 24px 16px 0;
  width: 120px;
  height: 31px;
  background: rgb(255 255 255 / 30%);
}

.ant-row-rtl #components-layout-demo-top .logo {
  float: right;
  margin: 16px 0 16px 24px;
}

[data-theme='dark'] .site-layout-content {
  background: #141414;
}

.my-card {
  display: flex;
  justify-content: space-between;

  .card-item {
    max-width: 300px;
    height: 300px;
    font-size: 32px;
    text-align: center;
    color: #fff;
    background-color: #333;
    flex: 1;
    line-height: 300px;
  }
}
</style>

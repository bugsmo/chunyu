<template>
  <n-tooltip trigger="hover">
    <template #trigger>
      <n-color-picker
        class="mr-16 h-40 w-80"
        :value="appStore.primaryColor"
        :swatches="primaryColors"
        :on-update:value="(v) => appStore.setPrimaryColor(v)"
      />
    </template>
    设置主题色
  </n-tooltip>
</template>

<script setup>
import { getPresetColors } from '@arco-design/color'
import { useAppStore } from '@/store'
const appStore = useAppStore()

const primaryColors = Object.entries(getPresetColors()).map(([, value]) => value.primary)
</script>

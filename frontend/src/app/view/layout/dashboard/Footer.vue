<template>
  <el-footer>&copy; {{ year }} {{ projectName }}.</el-footer>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { config } from '@/data/config/app'

export default defineComponent({
  name: 'Footer',
  setup: () => {
    return {
      projectName: config.projectName,
      year: new Date().getFullYear(),
    }
  },
})
</script>

<style lang="scss" scoped>
.el-footer {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>

<template>
  <main class="container">
    <router-view />
  </main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'AuthLayout',
})
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  max-width: 1320px;
  width: 100%;
  margin: 0 auto;
  padding: 0 10px;
}
</style>

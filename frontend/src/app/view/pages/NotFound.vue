<template>
  <main class="not-found min-h-100vh">
    <div class="box-content">
      <h1 class="title mb-2">404</h1>
      <router-link to="/" class="link-to-home">На главную</router-link>
    </div>
  </main>
</template>

<script lang="ts">
export default {
  name: 'NotFound',
}
</script>

<style lang="scss" scoped>
.not-found {
  display: flex;

  .box-content {
    text-align: center;
    margin: auto;
  }

  .link-to-home {
    color: #309ff6;
  }
}
</style>

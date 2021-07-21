<template>
  <el-card>
    <el-row type="flex" justify="center">
      <router-link to="/">
        <img src="@/app/assets/images/commerce/logo.svg" width="50" alt="" />
      </router-link>
    </el-row>
    <h1 class="title">{{ title }}</h1>
    <slot />
    <el-row type="flex" justify="center">
      <router-link :to="path" custom v-slot="{ navigate, href }">
        <el-link :href="href" @click="navigate">{{ pathTitle }}</el-link>
      </router-link>
    </el-row>
  </el-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'AuthCard',
  props: {
    title: {
      type: String,
      default: '',
    },
    path: {
      type: String,
      default: '/',
    },
    pathTitle: {
      type: String,
      default: 'Главная',
    },
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';
.title {
  margin-top: 0;
  margin-bottom: 20px;
  text-align: center;
}

.el-card {
  max-width: 420px;
  width: 100%;
  margin: auto;
}
</style>

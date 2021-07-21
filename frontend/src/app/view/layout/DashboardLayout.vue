<template>
  <el-container direction="vertical" class="min-h-100vh">
    <header-app />
    <router-view v-slot="{ Component }">
      <transition-page>
        <component :is="Component" />
      </transition-page>
    </router-view>
    <footer-app />
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import HeaderApp from '@/app/view/layout/dashboard/Header.vue'
import FooterApp from '@/app/view/layout/dashboard/Footer.vue'
import TransitionPage from '@/app/view/components/transition/TransitionPage.vue'

export default defineComponent({
  name: 'DashboardLayout',
  components: {
    TransitionPage,
    HeaderApp,
    FooterApp,
  },
})
</script>

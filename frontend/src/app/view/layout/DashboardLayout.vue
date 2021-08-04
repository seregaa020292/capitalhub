<template>
  <el-container direction="vertical" class="dashboard-container">
    <router-view v-slot="{ Component }">
      <transition-page>
        <component :is="Component" />
      </transition-page>
    </router-view>
  </el-container>
  <navbar-app />
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import NavbarApp from '@/app/view/layout/dashboard/Navbar.vue'
import TransitionPage from '@/app/view/components/transition/TransitionPage.vue'
import { ApplicationDataUseCaseDI } from '@/domain/application/module/di'

export default defineComponent({
  name: 'DashboardLayout',
  components: {
    TransitionPage,
    NavbarApp,
  },
  setup() {
    const applicationDataUseCase = ApplicationDataUseCaseDI()
    applicationDataUseCase.execute()
  },
})
</script>

<style lang="scss" scoped>
.dashboard-container {
  margin-bottom: 70px;
}
</style>

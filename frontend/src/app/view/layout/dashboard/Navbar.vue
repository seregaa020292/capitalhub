<template>
  <div class="navbar">
    <div class="menu-logo">
      <logo-company class="logo" />
    </div>
    <div class="menu-horizontal">
      <el-menu
        router
        :default-active="activeLink"
        mode="horizontal"
        class="menu"
        background-color="#1b2128"
        text-color="#fff"
        active-text-color="#409EFF"
      >
        <el-menu-item index="/dashboard" class="menu-item">
          <i class="el-icon-suitcase" />
          <span>Мой портфель</span>
        </el-menu-item>
        <el-menu-item index="/portfolios" class="menu-item">
          <i class="el-icon-menu" />
          <span>Мои портфели</span>
        </el-menu-item>
        <el-menu-item @click.once="onLogout">
          <i class="el-icon-unlock" />
          <span>Выйти</span>
        </el-menu-item>
      </el-menu>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useRoute } from 'vue-router'
import LogoCompany from '@/app/view/components/logoCompany/index.vue'
import { useLogout } from '@/app/hooks/common/useLogout'

export default defineComponent({
  name: 'Navbar',
  components: {
    LogoCompany,
  },
  setup() {
    const route = useRoute()
    const { onLogout } = useLogout()
    const activeLink = computed(() => route.matched[0].path)

    return {
      activeLink,
      onLogout,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/themes/element/variables';
@import '@/app/assets/styles/variables';

.navbar {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  background-color: $--color-main-darker;
  border-top: 1px solid $--color-main-dark-light;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  z-index: $--index-navbar;

  .menu-logo {
    display: flex;
    align-items: center;
    padding-top: 10px;
    padding-bottom: 10px;

    @include media-point('sm') {
      padding-right: 40px;
    }
  }

  .logo {
    height: 30px;
  }

  .menu-item {
    padding: 0 40px;
  }

  .menu-horizontal {
    padding-bottom: 10px;
    margin-bottom: -10px;
    overflow-x: auto;
  }

  .menu {
    display: flex;
    border-bottom: none;
  }
}
</style>

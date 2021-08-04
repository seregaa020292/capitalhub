<template>
  <el-header height="80px" class="header" id="topHeader">
    <div class="left-box">
      <logo-company class="logo" />
    </div>
    <div class="right-box">
      <el-link href="/logout" @click.once.prevent="onLogout" icon="el-icon-unlock">Выйти</el-link>
    </div>
  </el-header>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LogoCompany from '@/app/view/components/logoCompany/index.vue'
import { useLogout } from '@/app/hooks/useLogout'

export default defineComponent({
  name: 'Header',
  components: {
    LogoCompany,
  },
  setup: () => {
    const { onLogout } = useLogout()

    return {
      onLogout,
    }
  },
})
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-basis: 100%;
  padding-top: 20px;
  padding-bottom: 20px;
  margin-bottom: 20px;
  border-bottom: 1px solid #444;

  a {
    color: #fff;
  }

  & > .left-box {
    display: flex;
    align-items: center;

    .logo {
      height: 38px;
    }
  }
}
</style>

<template>
  <section class="header-container">
    <header class="header" id="topHeader">
      <div class="left-box">
        <logo-company native class="logo" />
      </div>
      <div class="right-box">
        <link-button-home type="primary" path="/dashboard">Мой портфель</link-button-home>
      </div>
    </header>
  </section>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import LogoCompany from '@/app/view/components/logoCompany/index.vue'
import LinkButtonHome from '@/app/view/containers/home/LinkButtonHome.vue'

export default defineComponent({
  name: 'Header',
  components: {
    LogoCompany,
    LinkButtonHome,
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';

.header-container {
  max-width: 1320px;
  width: 100%;
  margin: 0 auto;
  padding: 0 10px;

  .header {
    padding-top: 50px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    @include media-point('xs') {
      padding-top: 30px;
      justify-content: center;
    }

    & > .left-box {
      display: flex;
      align-items: center;

      .logo {
        height: 52px;
      }
    }

    & > .right-box {
      @include media-point('xs') {
        display: none;
      }
    }
  }
}
</style>

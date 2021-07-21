<template>
  <header-app />
  <main class="main-container">
    <div class="top-screen-info">
      <div class="left-box">
        <h1>{{ projectName }} - индивидуальный инструмент для управления портфелем активов</h1>
        <br />
        <div class="bot-btn-link box-row">
          <div class="box-col">
            <link-button-home type="primary" path="/dashboard">Мой портфель</link-button-home>
          </div>
          <div class="box-col">
            <link-button-home type="secondary" path="/register">Создать портфель</link-button-home>
          </div>
        </div>
      </div>
      <div class="right-box">
        <currency-box />
        <info-list />
      </div>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { CurrencyBox, InfoList } from '@/app/view/containers/home/components'
import LinkButtonHome from '@/app/view/containers/home/LinkButtonHome.vue'
import HeaderApp from '@/app/view/layout/home/Header.vue'
import { config } from '@/data/config/app'

export default defineComponent({
  name: 'Home',
  components: {
    InfoList,
    CurrencyBox,
    LinkButtonHome,
    HeaderApp,
  },
  setup: () => {
    return { projectName: config.projectName }
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';

.main-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: calc(100vh - 107px);
  max-width: 1320px;
  width: 100%;
  margin: 0 auto;
  padding: 0 10px;

  .top-screen-info {
    display: flex;
    flex-wrap: wrap;
    padding: 20px 0;
    width: 100%;

    & > .left-box {
      display: flex;
      flex: 1.4;
      flex-direction: column;

      @include media-point('sm') {
        margin-right: 35px;
      }

      h1 {
        margin-top: 0;
        margin-bottom: 15px;
        font-weight: bold;
        font-size: 58px;
        letter-spacing: 0.01em;

        @include media-point('xs') {
          font-size: 36px;
        }
      }

      & > .bot-btn-link {
        display: flex;
        flex-wrap: wrap;
        align-items: center;

        @include media-point('xs') {
          margin-bottom: 40px;
        }
      }

      .bot-btn-item {
        margin: 10px 0;

        @include media-point('xs') {
          width: 100%;
        }
      }

      .box-col {
        @include media-point('xs') {
          width: 100%;
          flex-basis: 100%;
        }
      }
    }

    & > .right-box {
      display: flex;
      flex-direction: column;
      flex: 1;
    }
  }
}
</style>

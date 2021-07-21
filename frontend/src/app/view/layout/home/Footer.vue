<template>
  <div class="footer-container">
    <footer>
      <div class="left-box">
        <logo-company class="logo" />
      </div>
      <div class="right-box">
        <link-button-home type="secondary" path="/register">Создать портфель</link-button-home>
      </div>
    </footer>
  </div>
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

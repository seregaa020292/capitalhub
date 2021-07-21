<template>
  <div id="app">
    <component :is="layout" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { useRoute } from 'vue-router'
import collects from '@/app/view/layout/collects'

export default defineComponent({
  name: 'App',
  setup() {
    const route = useRoute()
    const layout = computed(() => route.meta.layout || collects.defaultLayout)

    return {
      layout,
    }
  },
})
</script>

<style lang="scss" src="./assets/styles/index.scss" />

<template>
  <el-main>
    <h1 class="title text-center mb-3">{{ ticker }}</h1>
  </el-main>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'Ticker',
  setup: () => {
    const route = useRoute()

    return {
      ticker: route.params.ticker,
    }
  },
})
</script>

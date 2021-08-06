<template>
  <el-tooltip :content="title" placement="top-start">
    <router-link :to="link" class="ticker-link route-link white underline">
      <el-avatar
        :src="`https://eodhistoricaldata.com/img/logos/US/${ticker}.png`"
        :alt="ticker"
        shape="square"
        fit="cover"
      />
      {{ ticker }}
    </router-link>
  </el-tooltip>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'TickerView',
  props: {
    link: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
    ticker: {
      type: String,
      required: true,
    },
  },
})
</script>

<style lang="scss" scoped>
.ticker-link {
  display: inline-flex;
  flex-direction: column;
  align-items: center;
}

.ticker-logo {
  display: block;
  width: 50px;
}
</style>

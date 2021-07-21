<template>
  <span :class="`active-${num >= 0 ? 'pump' : 'dump'}`">
    <i :class="`el-icon-caret-${num >= 0 ? 'top' : 'bottom'}`" />
    <slot />
  </span>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'IndicatorCaretIcon',
  props: {
    num: {
      type: Number,
      default: 0,
    },
  },
})
</script>

<template>
  <span class="currency-caret">
    <i :class="{ pump: isPump, dump: !isPump }" />
  </span>
</template>

<script lang="ts">
export default {
  name: 'IndicatorCaret',
  props: {
    isPump: {
      type: Boolean,
      required: true,
    },
  },
}
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';

.currency-caret {
  display: inline-flex;
  flex-direction: column;
  align-items: center;
  height: 14px;
  width: 18px;
  vertical-align: middle;
  position: relative;

  .pump,
  .dump {
    width: 0;
    height: 0;
    border: solid 8px transparent;
    position: absolute;
  }

  .pump {
    border-bottom-color: $color-pump;
    top: -10px;
  }

  .dump {
    border-top-color: $color-dump;
    bottom: -2px;
  }
}
</style>

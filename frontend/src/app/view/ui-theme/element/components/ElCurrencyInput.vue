<template>
  <el-input ref="inputRef" v-model="formattedValue" />
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue'
import useCurrencyInput, { CurrencyInputOptions } from 'vue-currency-input'

export default defineComponent({
  name: 'ElCurrencyInput',
  props: {
    modelValue: {
      type: Number,
      default: null,
    },
    options: {
      type: Object as PropType<CurrencyInputOptions>,
      default: () => ({}),
    },
  },
  setup: (props) => {
    const { inputRef, formattedValue } = useCurrencyInput({
      locale: 'ru',
      hideCurrencySymbolOnFocus: false,
      hideGroupingSeparatorOnFocus: false,
      hideNegligibleDecimalDigitsOnFocus: false,
      valueRange: {
        min: 0,
      },
      ...props.options,
      currency: props.options.currency || 'RUB',
    })
    return { inputRef, formattedValue }
  },
})
</script>

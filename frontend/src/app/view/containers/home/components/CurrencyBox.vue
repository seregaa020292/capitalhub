<template>
  <div class="currency-box">
    <div class="item">
      <div class="left">
        <div class="img-box">
          <img src="@/app/assets/images/currency/bitcoin-color.svg" alt="Bitcoin" />
        </div>
        <div class="details">
          <div class="pre-name">BTC</div>
          <div class="name">Bitcoin</div>
        </div>
      </div>
      <div class="number">
        <indicator-caret :is-pump="isPumpBtc" />
        {{ currencyBtcRubFormat }} ₽
      </div>
    </div>
    <div class="item">
      <div class="left">
        <div class="img-box">
          <img src="@/app/assets/images/currency/dollar-color.svg" alt="Dollar" />
        </div>
        <div class="details">
          <div class="pre-name">USD</div>
          <div class="name">Dollar</div>
        </div>
      </div>
      <div class="number">
        <indicator-caret :is-pump="isPumpBtc" />
        {{ currencyUsdRubFormat }} ₽
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import IndicatorCaret from '@/app/view/components/currency/IndicatorCaret.vue'
import { useCurrencyBtc } from '@/app/hooks/useCurrencyBtc'

interface IProps {
  currencyUsdRub: number
}

export default defineComponent({
  name: 'CurrencyBox',
  components: {
    IndicatorCaret,
  },
  props: {
    currencyUsdRub: {
      type: Number,
      default: 73.56,
    },
  },
  setup: (props: IProps) => {
    const { state } = useCurrencyBtc()

    const currencyBtcRubFormat = computed(() =>
      Math.floor(state.currencyBtcUsd.currentPrice * props.currencyUsdRub).toLocaleString()
    )
    const currencyUsdRubFormat = computed(() => props.currencyUsdRub.toLocaleString())
    const isPumpBtc = computed(
      () => state.currencyBtcUsd.currentPrice - state.currencyBtcUsd.prevPrice > 0
    )

    return {
      currencyBtcRubFormat,
      currencyUsdRubFormat,
      isPumpBtc,
    }
  },
})
</script>

<style lang="scss" scoped>
@import '@/app/assets/styles/variables';

.currency-box {
  display: flex;
  z-index: 2;
  flex-direction: column;
  background: linear-gradient(180deg, rgba(44, 53, 62, 0.6) 0%, rgba(44, 53, 62, 0.414) 100%);
  border: 2px solid #313e47;
  backdrop-filter: blur(20px);
  border-radius: 36px;
  padding: 0 30px;
  max-width: 340px;
  width: 100%;

  @include media-point('xs') {
    min-width: 300px;
    padding: 0 15px;
  }

  .item {
    padding: 20px 0;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .item:not(:last-child) {
    border-bottom: 2px solid #313e47;
  }

  .item > .left {
    display: flex;
  }

  .item > .left > .img-box {
    width: 39px;
    height: 39px;
    margin-right: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .item > .left > .img-box img {
    width: 100%;
    height: 100%;
    object-fit: contain;
  }

  .item > .left > .details {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }

  .item > .left > .details > .pre-name {
    font-style: normal;
    font-weight: normal;
    font-size: 14px;
    color: #7c7d82;
  }

  .item > .left > .details > .name {
    font-style: normal;
    font-weight: 500;
    font-size: 16px;
  }

  .item > .number {
    font-style: normal;
    font-weight: 500;
    font-size: 22px;
  }
}
</style>

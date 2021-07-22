<template>
  <el-form
    :model="state.asset"
    :rules="rules"
    ref="ruleFormRef"
    novalidate
    @submit.prevent="onSubmit"
  >
    <el-form-item prop="marketId">
      <el-select
        v-model="state.asset.marketId"
        name="marketId"
        placeholder="Поиск по тикеру и названию"
        no-data-text="Не найдено..."
        no-match-text="Не найдено..."
        :remote-method="remoteMethod"
        :loading="state.market.loading"
        loading-text="Идет поиск"
        class="w-100"
        reserve-keyword
        remote
        clearable
        filterable
      >
        <template #prefix>
          <i class="el-input__icon el-icon-search"></i>
        </template>
        <el-option-group
          v-for="group in state.market.items"
          :key="group.label"
          :label="group.description"
        >
          <el-option
            v-for="item in group.options"
            :key="item.market_id"
            :label="`${item.ticker}: ${item.title}`"
            :value="item.market_id"
          />
        </el-option-group>
      </el-select>
    </el-form-item>
    <el-form-item prop="notationAt">
      <el-date-picker
        v-model="state.asset.notationAt"
        type="date"
        placeholder="Выберите дату"
        class="w-100"
      />
    </el-form-item>
    <el-row :gutter="8">
      <el-col :md="12">
        <el-form-item prop="amount">
          <el-currency-input name="amount" placeholder="Стоимость" v-model="state.asset.amount" />
        </el-form-item>
      </el-col>
      <el-col :md="12">
        <el-form-item prop="quantity">
          <el-input
            name="quantity"
            placeholder="Количество, шт."
            v-model.number="state.asset.quantity"
            class="w-100"
          />
        </el-form-item>
      </el-col>
    </el-row>
    <el-form-item><strong>Итоговая сумма</strong>: {{ totalAmount }}</el-form-item>
    <br />
    <el-button native-type="submit" type="primary" class="w-100">Добавить</el-button>
  </el-form>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, reactive, ref } from 'vue'
import { assetValidator } from '@/app/utils/validators'
import { IMarketOption } from '@/domain/market/entities/MarketEntity'
import {
  AssetAddUseCaseContainer,
  MarketSearchUseCaseContainer,
  PortfolioPresenterContainer,
} from '@/infrastructure/di/containers'
import { IAssetNotation } from '@/domain/asset/entities/AssetEntity'
import { currencyFormatter } from '@/utils/number'

export default defineComponent({
  name: 'AssetForm',
  setup: () => {
    const state = reactive({
      asset: {
        marketId: '',
        notationAt: '',
        amount: null,
        quantity: '',
      },
      market: {
        loading: false,
        items: [] as IMarketOption[],
      },
    })
    const marketSearchUseCase = MarketSearchUseCaseContainer()
    const assetAddUseCase = AssetAddUseCaseContainer()
    const portfolioPresenter = PortfolioPresenterContainer()

    const assetNotation: ComputedRef<IAssetNotation> = computed(() => ({
      ...state.asset,
      amount: Number(state.asset.amount),
      quantity: Number(state.asset.quantity),
      portfolioId: portfolioPresenter.portfolio().portfolioId,
    }))
    const totalAmount = computed(() =>
      currencyFormatter.rub.format(
        (assetNotation.value.amount || 0) * (assetNotation.value.quantity || 0)
      )
    )

    const remoteMethod = async (query: string) => {
      if (query.length >= 2) {
        state.market.loading = true
        const searchExecute = await marketSearchUseCase.execute(query)
        if (searchExecute) {
          state.market.items = searchExecute
        }
        state.market.loading = false
        return
      }
      state.market.items = []
    }

    const ruleFormRef: any = ref(null)
    const rules = ref({
      marketId: assetValidator.marketId,
      notationAt: assetValidator.notationAt,
      amount: assetValidator.amount,
      quantity: assetValidator.quantity,
    })

    const onSubmit = () => {
      ruleFormRef.value.validate(async (valid: boolean) => {
        if (!valid) return false
        await assetAddUseCase.execute(assetNotation.value)
        ruleFormRef.value.resetFields()
      })
    }

    return {
      state,
      totalAmount,
      rules,
      ruleFormRef,
      onSubmit,
      remoteMethod,
    }
  },
})
</script>

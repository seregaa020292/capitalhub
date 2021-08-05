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
            :key="item.marketId"
            :label="`${item.ticker}: ${item.title}`"
            :value="item.marketId"
          />
        </el-option-group>
      </el-select>
    </el-form-item>
    <el-row :gutter="8">
      <el-col :md="12">
        <el-form-item prop="notationAt">
          <el-date-picker
            v-model="state.asset.notationAt"
            type="date"
            placeholder="Выберите дату"
            class="w-100"
          />
        </el-form-item>
      </el-col>
      <el-col :md="12">
        <el-select v-model="state.operation" placeholder="Операция" class="w-100">
          <el-option
            v-for="({ operation, label }, idx) in state.operations"
            :key="idx"
            :label="label"
            :value="operation"
          />
        </el-select>
      </el-col>
    </el-row>
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
    <el-row :gutter="8">
      <el-col :md="12">
        <el-button native-type="submit" type="primary" class="w-100">
          Добавить и продолжить
        </el-button>
      </el-col>
      <el-col :md="12">
        <el-button @click="onSubmitClose" type="info" class="w-100">Добавить и закрыть</el-button>
      </el-col>
    </el-row>
  </el-form>
</template>

<script lang="ts">
import { computed, ComputedRef, defineComponent, reactive, ref } from 'vue'
import { assetValidator } from '@/app/utils/validators'
import { IMarketOption } from '@/domain/market/entities/MarketEntity'
import { AssetAddUseCaseDI } from '@/domain/asset/module/di'
import { MarketSearchUseCaseDI } from '@/domain/market/module/di'
import { PortfolioPresenterDI } from '@/domain/portfolio/module/di'
import { IAssetNotation } from '@/domain/asset/entities/AssetEntity'
import { currencyFormatter } from '@/utils/number'

export default defineComponent({
  name: 'AssetForm',
  emits: ['close-form'],
  setup: (props, { emit }) => {
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
      operation: 1,
      operations: [
        {
          operation: 1,
          label: 'Покупка',
        },
        {
          operation: -1,
          label: 'Продажа',
        },
      ],
    })
    const marketSearchUseCase = MarketSearchUseCaseDI()
    const assetAddUseCase = AssetAddUseCaseDI()
    const portfolioPresenter = PortfolioPresenterDI()

    const assetNotation: ComputedRef<IAssetNotation> = computed(() => ({
      ...state.asset,
      amount: Number(state.asset.amount) * state.operation,
      quantity: Number(state.asset.quantity),
      portfolioId: portfolioPresenter.portfolio().portfolioId,
    }))
    const totalAmount = computed(() =>
      currencyFormatter.rub.format(
        (assetNotation.value.amount || 0) * (assetNotation.value.quantity || 0)
      )
    )

    const remoteMethod = async (query: string) => {
      if (query.length < 2) {
        state.market.items = []
        return
      }
      state.market.loading = true
      const searchExecute = await marketSearchUseCase.execute(query)
      if (searchExecute) {
        state.market.items = searchExecute
      }
      state.market.loading = false
    }

    const ruleFormRef: any = ref(null)
    const rules = ref({
      marketId: assetValidator.marketId,
      notationAt: assetValidator.notationAt,
      amount: assetValidator.amount,
      quantity: assetValidator.quantity,
    })

    const assetAdded = (fn: (hasAdded: boolean) => void) => {
      ruleFormRef.value.validate(async (valid: boolean) => {
        if (!valid) {
          return false
        }
        const hasAdded = await assetAddUseCase.execute(assetNotation.value)
        fn(hasAdded)
      })
    }

    const onSubmit = () => assetAdded((hasAdded) => hasAdded && ruleFormRef.value.resetFields())
    const onSubmitClose = () => assetAdded((hasAdded) => hasAdded && emit('close-form'))

    return {
      state,
      totalAmount,
      rules,
      ruleFormRef,
      onSubmit,
      onSubmitClose,
      remoteMethod,
    }
  },
})
</script>

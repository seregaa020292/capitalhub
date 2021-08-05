<template>
  <el-form
    :model="state.portfolio"
    :rules="rules"
    ref="ruleFormRef"
    novalidate
    @submit.prevent="onSubmit"
  >
    <el-form-item prop="title">
      <el-input v-model="state.portfolio.title" name="title" placeholder="Название портфеля" />
    </el-form-item>
    <el-form-item prop="currencyId">
      <el-select v-model="state.portfolio.currencyId" placeholder="Портфель в валюте" class="w-100">
        <el-option
          v-for="{ description, currencyId } in currencies"
          :key="currencyId"
          :label="description"
          :value="currencyId"
        />
      </el-select>
    </el-form-item>
    <el-row>
      <el-col :offset="6" :md="12">
        <el-button native-type="submit" type="primary" class="w-100">
          {{ state.isEditing ? 'Сохранить' : 'Создать' }}
        </el-button>
      </el-col>
    </el-row>
  </el-form>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, computed, PropType } from 'vue'
import { portfolioValidator } from '@/app/utils/validators'
import { PortfolioAddUseCaseDI, PortfolioEditUseCaseDI } from '@/domain/portfolio/module/di'
import { ApplicationPresenterDI } from '@/domain/application/module/di'
import {
  IPortfolioChangeFields,
  IPortfolioEditFields,
} from '@/domain/portfolio/entities/PortfolioEntity'

export default defineComponent({
  name: 'Form',
  props: {
    portfolioEdit: {
      type: Object as PropType<IPortfolioEditFields>,
      default: (): IPortfolioEditFields => ({
        portfolioId: '',
        currencyId: '',
        title: '',
      }),
    },
  },
  setup(props) {
    const state = reactive({
      isEditing: !!props.portfolioEdit.portfolioId,
      portfolio: {
        title: props.portfolioEdit.title,
        currencyId: props.portfolioEdit.currencyId,
      } as IPortfolioChangeFields,
    })
    const ruleFormRef: any = ref(null)
    const rules = ref({
      title: portfolioValidator.title,
      currencyId: portfolioValidator.currencyId,
    })
    const portfolioAddUseCase = PortfolioAddUseCaseDI()
    const portfolioEditUseCase = PortfolioEditUseCaseDI()
    const applicationPresenter = ApplicationPresenterDI()
    const currencies = computed(() => applicationPresenter.getDashboard().currencies)

    const onSubmit = () => {
      ruleFormRef.value.validate(async (valid: boolean) => {
        if (!valid) {
          return false
        }

        if (state.isEditing) {
          await portfolioEditUseCase.execute({
            portfolioId: props.portfolioEdit.portfolioId,
            ...state.portfolio,
          })
        } else {
          await portfolioAddUseCase.execute(state.portfolio)
        }

        ruleFormRef.value.resetFields()
      })
    }

    return {
      state,
      currencies,
      rules,
      ruleFormRef,
      onSubmit,
    }
  },
})
</script>

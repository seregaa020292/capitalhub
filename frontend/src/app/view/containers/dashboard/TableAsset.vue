<template>
  <el-table v-loading="loading" :data="assets" empty-text="Нет данных">
    <el-table-column fixed type="index" label="#" />
    <el-table-column fixed width="125px" label="Наименование">
      <template v-slot="{ row }">
        <el-tooltip :content="row.title" placement="top-start">
          <router-link
            :to="`/dashboard/${row.ticker}`"
            class="ticker-link route-link white underline"
          >
            <el-avatar
              :src="`https://eodhistoricaldata.com/img/logos/US/${row.ticker}.png`"
              :alt="row.ticker"
              shape="square"
              fit="cover"
            />
            {{ row.ticker }}
          </router-link>
        </el-tooltip>
      </template>
    </el-table-column>
    <el-table-column width="140px" sortable prop="currentPrice" label="Текущая цена">
      <template v-slot="{ row }">{{ currencyFormatter.rub.format(row.currentPrice) }}</template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="currentValue" label="Текущая стоимость">
      <template v-slot="{ row }">
        {{ currencyFormatter.rub.format(row.currentPrice * row.totalQuantity) }}
      </template>
    </el-table-column>
    <el-table-column width="120px" sortable prop="totalQuantity" label="Кол-во, шт" />
    <el-table-column
      width="200px"
      sortable
      prop="averagePurchasePrice"
      label="Средняя цена покупки"
    >
      <template v-slot="{ row }">
        {{ currencyFormatter.rub.format(row.averagePurchasePrice) }}
      </template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="totalAmount" label="Стоимость покупки">
      <template v-slot="{ row }">{{ currencyFormatter.rub.format(row.totalAmount) }}</template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="changeTotalPercent" label="Изменение, всего">
      <template v-slot="{ row }">
        <indicator-caret-icon
          :num="(row.currentPrice * row.totalQuantity - row.totalAmount) / row.totalAmount"
        >
          {{
            ((row.currentPrice * row.totalQuantity - row.totalAmount) / row.totalAmount).toFixed(2)
          }}%
        </indicator-caret-icon>
      </template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="changePerDayPercent" label="Изменение, за день">
      <template v-slot="{ row }">
        <span class="active-dump">
          <i class="el-icon-caret-bottom" />
          {{ row.changePerDayPercent }}%
        </span>
      </template>
    </el-table-column>
    <el-table-column label="Информация">
      <template v-slot="{ row }">
        <el-tooltip :content="tooltipNotation(row.firstNotationAt)" placement="top-start">
          <i class="el-icon-time" />
        </el-tooltip>
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts">
import { computed, defineComponent } from 'vue'
import { dayjs } from '@/utils/dayjs'
import { currencyFormatter } from '@/utils/number'
import { AssetPresenterDI } from '@/domain/asset/module/di'
import IndicatorCaretIcon from '@/app/view/components/currency/IndicatorCaretIcon.vue'

export default defineComponent({
  name: 'TableAsset',
  components: {
    IndicatorCaretIcon,
  },
  setup: () => {
    const assetPresenter = AssetPresenterDI()
    const tooltipNotation = (notationAt: string) => `Актив добавлен: ${dayjs().to(notationAt)}`

    const assets = computed(() => assetPresenter.assets())
    const loading = computed(() => assetPresenter.loadingAssets())

    return {
      assets,
      loading,
      tooltipNotation,
      currencyFormatter,
    }
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

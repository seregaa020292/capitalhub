<template>
  <el-table v-loading="loading" :data="assets" empty-text="Нет данных">
    <el-table-column fixed type="index" label="#" />
    <el-table-column fixed width="125px" label="Наименование">
      <template v-slot="{ row }">
        <ticker-view
          v-if="row.ticker !== undefined"
          :link="`/dashboard/${row.ticker}`"
          :title="row.title"
          :ticker="row.ticker"
        />
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
    <el-table-column width="200px" sortable prop="averageTotalAmount" label="Средняя цена покупки">
      <template v-slot="{ row }">
        {{ currencyFormatter.rub.format(row.averageTotalAmount) }}
      </template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="totalAmount" label="Стоимость покупки">
      <template v-slot="{ row }">
        {{ currencyFormatter.rub.format(row.totalAmount) }}
      </template>
    </el-table-column>
    <el-table-column width="175px" sortable prop="changeTotalPercent" label="Изменение, всего">
      <template v-slot="{ row }">
        <indicator-caret-icon :num="row.currentPrice / row.averageTotalAmount">
          {{ (row.currentPrice / row.averageTotalAmount).toFixed(2) }}%
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
import TickerView from '@/app/view/components/ticker/TickerView.vue'

export default defineComponent({
  name: 'AssetTable',
  components: {
    IndicatorCaretIcon,
    TickerView,
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

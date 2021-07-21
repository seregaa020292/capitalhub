import { createApp } from 'vue'
import {
  ElButton,
  ElButtonGroup,
  ElCard,
  ElCol,
  ElCollapseTransition,
  ElDialog,
  ElDropdown,
  ElDropdownItem,
  ElDropdownMenu,
  ElForm,
  ElFormItem,
  ElIcon,
  ElInput,
  ElInputNumber,
  ElSelect,
  ElOption,
  ElOptionGroup,
  ElLoading,
  ElMenu,
  ElMenuItem,
  ElContainer,
  ElHeader,
  ElMain,
  ElFooter,
  ElTable,
  ElTableColumn,
  ElMessage,
  ElRow,
  ElTag,
  ElTimeline,
  ElTimelineItem,
  ElCheckbox,
  ElLink,
  ElDatePicker,
  ElAvatar,
  ElImage,
  ElTooltip,
  locale
} from 'element-plus'
import { ElCurrencyInput } from '@/app/view/ui-theme/element/components'

import lang from 'element-plus/lib/locale/lang/ru'
import 'dayjs/locale/ru'
import './index.scss'

export default (app: ReturnType<typeof createApp>): void => {
  locale(lang)

  const components = [
    ElButton,
    ElButtonGroup,
    ElDialog,
    ElForm,
    ElFormItem,
    ElInput,
    ElInputNumber,
    ElCurrencyInput,
    ElSelect,
    ElOption,
    ElOptionGroup,
    ElMessage,
    ElMenu,
    ElMenuItem,
    ElContainer,
    ElHeader,
    ElMain,
    ElFooter,
    ElTable,
    ElTableColumn,
    ElRow,
    ElCol,
    ElDropdownMenu,
    ElTimeline,
    ElTimelineItem,
    ElDropdownItem,
    ElDropdown,
    ElCard,
    ElTag,
    ElIcon,
    ElCollapseTransition,
    ElCheckbox,
    ElLink,
    ElDatePicker,
    ElAvatar,
    ElImage,
    ElTooltip,
  ]

  components.forEach((component) => {
    app.component(component.name, component)
  })

  app.config.globalProperties.$loading = ElLoading.service
}

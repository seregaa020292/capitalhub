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
  ElPopconfirm,
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
  locale,
} from 'element-plus'
import { ElCurrencyInput } from '@/app/themes/element/components'
import { ConfirmService } from '@/services/message/ConfirmService'

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
    ElPopconfirm,
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

  const plugins = [
    ElLoading,
  ]

  components.forEach((component) => {
    app.component(component.name, component)
  })

  plugins.forEach(plugin => {
    app.use(plugin)
  })

  app.provide('$loading', ElLoading.service)
  app.provide('$confirm', ConfirmService.instance())
}

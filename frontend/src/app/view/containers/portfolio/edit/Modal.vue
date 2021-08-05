<template>
  <el-dialog
    v-model="dialogVisible"
    :title="titleModal"
    custom-class="dialog-size small"
    destroy-on-close
    append-to-body
  >
    <edit-form :portfolio-edit="portfolioEdit" />
  </el-dialog>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
import EditForm from '@/app/view/containers/portfolio/edit/Form.vue'
import { useModalHandleInject } from '@/app/hooks/useModalHandleProvideInject'
import { usePortfolioEditInject } from '@/app/hooks/usePortfolioEditProvideInject'

export default defineComponent({
  name: 'Modal',
  components: {
    EditForm,
  },
  setup() {
    const { dialogVisible } = useModalHandleInject('portfolio')
    const { portfolioEdit } = usePortfolioEditInject()
    const titleModal = computed(() =>
      portfolioEdit.value !== undefined ? 'Редактирование портфеля' : 'Новый портфель'
    )

    return {
      titleModal,
      dialogVisible,
      portfolioEdit,
    }
  },
})
</script>

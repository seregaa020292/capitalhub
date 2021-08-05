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
import { usePortfolioModalInject } from '@/app/hooks/portfolio/usePortfolioModalProvideInject'
import { usePortfolioEditInject } from '@/app/hooks/portfolio/usePortfolioEditProvideInject'

export default defineComponent({
  name: 'Modal',
  components: {
    EditForm,
  },
  setup() {
    const { dialogVisible } = usePortfolioModalInject()
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

<template>
  <el-button type="primary" @click="dialogOpenHandle">
    <i class="el-icon-circle-plus" />
    Добавить еще портфель
  </el-button>
  <el-dialog
    v-model="dialogVisible"
    title="Новый портфель"
    custom-class="dialog-size small"
    destroy-on-close
    append-to-body
  >
    <edit-form />
  </el-dialog>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import EditForm from '@/app/view/containers/portfolio/edit/Form.vue'

export default defineComponent({
  name: 'Modal',
  components: {
    EditForm,
  },
  setup() {
    const dialogVisible = ref(false)

    const dialogOpenHandle = () => {
      dialogVisible.value = true
    }

    const dialogClosedHandle = () => {
      dialogVisible.value = false
    }

    return {
      dialogVisible,
      dialogOpenHandle,
      dialogClosedHandle,
    }
  },
})
</script>

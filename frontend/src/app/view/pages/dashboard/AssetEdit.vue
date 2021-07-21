<template>
  <el-dialog
    v-model="dialogVisible"
    title="Добавить в портфель актив"
    custom-class="dialog-size small"
    width="100%"
    :fullscreen="false"
    append-to-body
    @closed="dialogClosedHandle"
  >
    <asset-form />
  </el-dialog>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import AssetForm from '@/app/view/containers/dashboard/AssetForm.vue'

export default defineComponent({
  name: 'AssetEdit',
  components: {
    AssetForm,
  },
  setup: () => {
    const dialogVisible = ref(true)
    const router = useRouter()

    const dialogClosedHandle = () => {
      router.push({ name: 'dashboard' })
    }

    return {
      dialogVisible,
      dialogClosedHandle,
    }
  },
})
</script>

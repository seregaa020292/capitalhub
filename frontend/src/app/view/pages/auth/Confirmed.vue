<template>
  <div class="confirmed text-center">
    <h1 class="mb-2">Ждите. Идет подтверждение почты...</h1>
    <router-link to="/" custom v-slot="{ navigate, href }">
      <el-link :href="href" @click="navigate" type="primary">На главную</el-link>
    </router-link>
  </div>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { AuthConfirmedUseCaseDI } from '@/domain/auth/module/di'
import { MessageService } from '@/services/message/MessageService'

export default defineComponent({
  name: 'Confirmed',
  setup: () => {
    const { params } = useRoute()
    const router = useRouter()
    const authConfirmedUseCase = AuthConfirmedUseCaseDI()

    onBeforeMount(async () => {
      if (params.code !== '') {
        const isConfirmed = await authConfirmedUseCase.execute(params.code as string)

        if (isConfirmed) {
          MessageService.instance().success('Почта подтверждена.')
          router.replace({ name: 'dashboard' })
          return
        }
      }

      MessageService.instance().error('Не удалось подтвердить почту.')
      router.replace('/')
    })
  },
})
</script>

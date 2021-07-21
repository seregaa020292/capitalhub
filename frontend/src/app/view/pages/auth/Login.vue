<template>
  <auth-card title="Вход" path="/register" path-title="Регистрация">
    <el-form
      :model="credentials"
      :rules="rules"
      ref="ruleFormRef"
      novalidate
      @submit.prevent="onSubmit"
    >
      <el-form-item prop="email">
        <el-input type="email" name="email" placeholder="Email*" v-model="credentials.email" />
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          name="password"
          placeholder="Пароль*"
          v-model="credentials.password"
          show-password
        />
      </el-form-item>
      <br />
      <el-row type="flex" justify="center">
        <el-form-item class="w-100">
          <el-button native-type="submit" type="primary" class="w-100">Войти</el-button>
        </el-form-item>
      </el-row>
    </el-form>
  </auth-card>
</template>

<script lang="ts">
import { defineComponent, Ref, ref } from 'vue'
import { useRouter } from 'vue-router'
import { authValidator } from '@/app/utils/validators'
import AuthCard from '@/app/view/containers/auth/AuthCard.vue'
import { AuthServiceContainer } from '@/infrastructure/di/containers'

export default defineComponent({
  name: 'Login',
  components: {
    AuthCard,
  },
  setup: () => {
    const credentials = ref({
      email: '',
      password: '',
    })

    const ruleFormRef: Ref = ref(null)

    const router = useRouter()

    const rules = ref({
      email: authValidator.email,
      password: authValidator.password,
    })

    const onSubmit = () => {
      ruleFormRef.value.validate(async (valid: boolean) => {
        if (!valid) return false

        const isLogin = await AuthServiceContainer().login({
          email: credentials.value.email,
          password: credentials.value.password,
        })

        if (isLogin) {
          ruleFormRef.value.resetFields()
          router.push({ name: 'dashboard' })
        }
      })
    }

    return {
      credentials,
      rules,
      ruleFormRef,
      onSubmit,
    }
  },
})
</script>

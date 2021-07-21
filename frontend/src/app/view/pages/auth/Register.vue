<template>
  <auth-card title="Регистрация" path="/login" path-title="Войти">
    <el-form
      :model="candidate"
      :rules="rules"
      ref="ruleFormRef"
      novalidate
      @submit.prevent="onSubmit"
    >
      <el-form-item prop="email">
        <el-input type="email" name="email" placeholder="Email" v-model="candidate.email" />
      </el-form-item>
      <el-form-item prop="name">
        <el-input type="text" name="name" placeholder="Имя" v-model="candidate.name" />
      </el-form-item>
      <el-form-item prop="password">
        <el-input
          type="password"
          name="password"
          placeholder="Пароль"
          v-model="candidate.password"
          show-password
        />
      </el-form-item>
      <el-form-item prop="agree">
        <el-checkbox name="agree" v-model="candidate.agree">
          Я согласен с политикой конфиденциальности
        </el-checkbox>
      </el-form-item>
      <el-row type="flex" justify="center" align="middle">
        <el-form-item class="w-100">
          <el-button native-type="submit" type="primary" class="w-100">Создать</el-button>
        </el-form-item>
      </el-row>
    </el-form>
  </auth-card>
</template>

<script lang="ts">
import { defineComponent, ref, Ref } from 'vue'
import { useRouter } from 'vue-router'
import { authValidator } from '@/app/utils/validators'
import AuthCard from '@/app/view/containers/auth/AuthCard.vue'
import { AuthServiceContainer } from '@/infrastructure/di/containers'

export default defineComponent({
  name: 'Register',
  components: {
    AuthCard,
  },
  setup: () => {
    const candidate = ref({
      name: '',
      email: '',
      password: '',
      agree: true,
    })

    const router = useRouter()

    const ruleFormRef: Ref = ref(null)

    const rules = ref({
      email: authValidator.email,
      name: authValidator.name,
      password: authValidator.password,
      agree: authValidator.agree,
    })

    const onSubmit = () => {
      ruleFormRef.value.validate(async (valid: boolean) => {
        if (!valid) return false

        const isReg = await AuthServiceContainer().registration({
          email: candidate.value.email,
          name: candidate.value.name,
          password: candidate.value.password,
        })

        if (isReg) {
          ruleFormRef.value.resetFields()
          router.push({ name: 'login' })
        }
      })
    }

    return {
      candidate,
      rules,
      ruleFormRef,
      onSubmit,
    }
  },
})
</script>

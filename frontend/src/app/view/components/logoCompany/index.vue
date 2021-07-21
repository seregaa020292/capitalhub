<template>
  <a href="/" @click="linkHandle" class="logo-box">
    <img src="@/app/assets/images/commerce/logo-landscape.svg" class="logo" alt="Логотип" />
  </a>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: 'Logo',
  props: {
    native: {
      type: Boolean,
      default: false,
    },
  },
  setup: (props) => {
    const router = useRouter()

    const linkHandle = (event: MouseEvent) => {
      if (!props.native) {
        event.preventDefault()
        router.push('/')
      }
    }

    return {
      linkHandle,
    }
  },
})
</script>

<style lang="scss" scoped>
.logo {
  height: inherit;
  width: inherit;
}
</style>

<template>
  <transition :name="state.transitionName" mode="out-in" appear>
    <slot />
  </transition>
</template>

<script lang="ts">
import { defineComponent, reactive } from 'vue'
import { isMobile } from '@/utils/device'

const DEFAULT_TRANSITION = 'route'

export default defineComponent({
  name: 'TransitionPage',
  setup: () => {
    const state = reactive({
      transitionName: DEFAULT_TRANSITION,
    })

    if (isMobile()) {
      state.transitionName = ''
    }

    return {
      state,
    }
  },
})
</script>

<style lang="scss" scoped>
/* route transitions */
.route-enter-from {
  opacity: 0;
  transform: translateX(100px);
}
.route-enter-active {
  transition: all 0.3s ease-out;
}
.route-leave-to {
  opacity: 0;
  transform: translateX(-100px);
}
.route-leave-active {
  transition: all 0.3s ease-in;
}
</style>

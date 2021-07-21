<template>
  <router-link :to="path" class="bot-btn-item" :class="typeClass"><slot /></router-link>
</template>

<script lang="ts">
import { computed, defineComponent, PropType } from 'vue'

interface BtnType {
  primary: string
  secondary: string
}

interface IProps {
  type: BtnType
  path: string
}

const btnTypes: BtnType = {
  primary: 'btn-primary-large',
  secondary: 'btn-secondary-large',
}

export default defineComponent({
  name: 'LinkButtonHome',
  props: {
    type: {
      type: String as keyof PropType<BtnType>,
      default: 'primary',
      validator: (value: string) => Object.keys(btnTypes).includes(value),
    },
    path: {
      type: String,
      required: true,
    },
  },
  setup: (props: IProps) => {
    const typeClass = computed(() => btnTypes[props.type as keyof PropType<BtnType>])

    return {
      typeClass,
    }
  },
})
</script>

<style lang="scss" scoped>
.btn-primary-large,
.btn-secondary-large {
  display: inline-block;
  font-style: normal;
  font-size: 20px;
  border-radius: 16px;
  text-align: center;
  transition: background 0.3s ease;
  border: none;
  background: none;
}

.btn-primary-large {
  font-weight: bold;
  position: relative;
  color: #ffffff;
  padding: 17px 39px;
  background: #1d75e3 linear-gradient(180deg, #309ff6 0%, #1d75e3 100%);

  &:hover {
    background: #1e252b;
  }
}

.btn-secondary-large {
  position: relative;
  font-weight: 500;
  padding: 18px 39px;
  background: #1e252b;
  color: #ffffff;
  border: 0 solid rgba(255, 255, 255, 0.4);

  &:hover {
    background: lighten(#1e252b, 7%);
  }
}
</style>

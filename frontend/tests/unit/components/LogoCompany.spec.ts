import { shallowMount } from '@vue/test-utils'
import LogoCompany from '@/app/view/components/logoCompany/index.vue'

describe('LogoCompany.vue', () => {
  it('renders props.native when passed', () => {
    const native = true
    const wrapper = shallowMount(LogoCompany, {
      props: { native },
    })
    expect(wrapper.props().native).toEqual(native)
  })
})

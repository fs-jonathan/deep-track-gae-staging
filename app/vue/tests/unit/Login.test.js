import { shallowMount } from '@vue/test-utils'
import Login from '@/components/Login.vue'

describe('Login.vue', () => {
  it('renders correctly', () => {
    const msg = 'new message'
    const wrapper = shallowMount(Login, {
      propsData: { result: {
        msg: msg
      } }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})

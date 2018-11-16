import { shallowMount } from '@vue/test-utils'
import BaseRow from '@/components/BaseRow.vue'

describe('HelloWorld.vue', () => {
  it('renders correctly', () => {
    const msg = 'new message'
    const wrapper = shallowMount(BaseRow, {
      propsData: { result: {
        title: msg,
        cost: 0,
        compare: 1,
        rate: 1,
        id: 2
      } }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})

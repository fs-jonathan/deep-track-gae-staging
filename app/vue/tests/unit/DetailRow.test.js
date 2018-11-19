import { shallowMount } from '@vue/test-utils'
import DetailRow from '@/components/DetailRow.vue'

describe('DetailRow.vue', () => {
  it('renders correctly', () => {
    const msg = 'new message'
    const wrapper = shallowMount(DetailRow, {
      propsData: { result: {
        title: msg,
        cost: 0,
        id: 2
      } }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
    expect(wrapper.element).toMatchSnapshot()
  })
})

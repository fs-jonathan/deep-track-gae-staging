import { shallowMount, createLocalVue } from '@vue/test-utils'
import GetRecord from '@/components/GetRecord.vue'
import VueElementLoading from 'vue-element-loading'
import VCalendar from 'v-calendar'
import BootstrapVue from 'bootstrap-vue'
import VueMoment from 'vue-moment'

const localVue = createLocalVue()
localVue.use(VCalendar)
localVue.use(BootstrapVue)
localVue.component('VueElementLoading', VueElementLoading)

describe('GetRecord.vue', () => {
  it('renders correctly', () => {
    const msg = 'new message'
    const wrapper = shallowMount(GetRecord, {
      propsData: { result: {
        title: msg,
        cost: 0,
        id: 2
      } },
      localVue,
      stubs: {
        BField: true
      }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})

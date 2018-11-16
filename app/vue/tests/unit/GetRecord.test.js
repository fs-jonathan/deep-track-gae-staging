import { shallowMount, createLocalVue } from '@vue/test-utils'
import GetRecord from '@/components/GetRecord.vue'
import VueElementLoading from 'vue-element-loading'
import VCalendar from 'v-calendar'
import BootstrapVue from 'bootstrap-vue'

const localVue = createLocalVue()
localVue.use(VCalendar)
localVue.use(BootstrapVue)
localVue.component('VueElementLoading', VueElementLoading)

describe('GetRecord.vue', () => {
  it('renders correctly', async () => {
    const msg = 'new message'
    const wrapper = shallowMount(GetRecord, {
      attachToDocument: true,
      propsData: { result: {
        title: msg,
        cost: 0,
        id: 2
      } },
      localVue,
      stubs: {
        BField: true
      },
      mocks: {
        $axios: {
          post: () => Promise.resolve({ data: 'value' })
        }
      }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})

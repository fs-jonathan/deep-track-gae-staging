import { shallowMount, createLocalVue } from '@vue/test-utils'
import GetRecord from '@/components/GetRecord.vue'
import VueElementLoading from 'vue-element-loading'
import VCalendar from 'v-calendar'
import BootstrapVue from 'bootstrap-vue'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()
localVue.use(VCalendar)
localVue.use(BootstrapVue)
localVue.component('VueElementLoading', VueElementLoading)

const data = [{"id":1,"cost":82,"compare":0,"rate":0},{"id":2,"cost":14,"compare":0.1,"rate":0.1},{"id":3,"cost":1517,"compare":0,"rate":0},{"id":4,"cost":0,"compare":0,"rate":0},{"id":5,"cost":1517,"compare":0,"rate":0}]

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
          post: () => Promise.resolve({ data: data })
        }
      }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
    await flushPromises()
    expect(wrapper.element).toMatchSnapshot()
  })
})

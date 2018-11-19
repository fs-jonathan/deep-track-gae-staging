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

const data = [{"id":1,"title":"本日","subtitle":"（現時点まで）","cost":82,"compare":0,"rate":0},{"id":2,"title":"昨日","subtitle":"先週と同じ曜日との比較","cost":14,"compare":0.1,"rate":0.1},{"id":3,"title":"今月（現時点まで）","subtitle":"先月の同じ日との比較","cost":1517,"compare":0,"rate":0},{"id":4,"title":"今月（現時点まで）","subtitle":"先月の同じ日との比較","cost":0,"compare":0,"rate":0},{"id":5,"title":"全期間","subtitle":"","cost":1517,"compare":0,"rate":0}]

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

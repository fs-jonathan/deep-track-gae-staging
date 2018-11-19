import { shallowMount, createLocalVue } from '@vue/test-utils'
import VueRouter from 'vue-router'
import ShowRecord from '@/components/ShowRecord.vue'
import VueElementLoading from 'vue-element-loading'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()
localVue.use(VueRouter)
localVue.component('VueElementLoading', VueElementLoading)

const router = new VueRouter()

const data = [{"id":6,"title":"見積もり収益額","subtitle":"","cost":9,"compare":0,"rate":0},{"id":7,"title":"ページビュー","subtitle":"","cost":193,"compare":0,"rate":0},{"id":8,"title":"表示回数","subtitle":"","cost":8,"compare":0,"rate":0},{"id":9,"title":"ページCTR","subtitle":"","cost":0.7269779556872005,"compare":0,"rate":0},{"id":10,"title":"クリック率","subtitle":"","cost":0.8896918783667795,"compare":0,"rate":0},{"id":11,"title":"カバレッジ","subtitle":"","cost":0.5157602624061407,"compare":0,"rate":0}]

describe('ShowRecord.vue', () => {
  it('renders correctly', async () => {
    const wrapper = shallowMount(ShowRecord, {
      localVue,
      router,
      attachToDocument: true,
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

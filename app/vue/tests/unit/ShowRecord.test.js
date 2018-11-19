import { shallowMount, createLocalVue } from '@vue/test-utils'
import VueRouter from 'vue-router'
import ShowRecord from '@/components/ShowRecord.vue'
import VueElementLoading from 'vue-element-loading'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()
localVue.use(VueRouter)
localVue.component('VueElementLoading', VueElementLoading)

const router = new VueRouter()

const data = [{"id":6,"cost":9,"compare":0,"rate":0},{"id":7,"cost":193,"compare":0,"rate":0},{"id":8,"cost":8,"compare":0,"rate":0},{"id":9,"cost":0.7269779556872005,"compare":0,"rate":0},{"id":10,"cost":0.8896918783667795,"compare":0,"rate":0},{"id":11,"cost":0.5157602624061407,"compare":0,"rate":0}]

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

import { shallowMount, createLocalVue } from '@vue/test-utils'
import VueRouter from 'vue-router'
import ShowRecord from '@/components/ShowRecord.vue'
import VueElementLoading from 'vue-element-loading'
import flushPromises from 'flush-promises'

const localVue = createLocalVue()
localVue.use(VueRouter)
localVue.component('VueElementLoading', VueElementLoading)

const router = new VueRouter()

describe('ShowRecord.vue', () => {
  it('renders correctly', async () => {
    const wrapper = shallowMount(ShowRecord, {
      localVue,
      router,
      attachToDocument: true,
      mocks: {
        $axios: {
          post: () => Promise.resolve({ data: 'value' })
        }
      }
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
    await flushPromises()
    expect(wrapper.vm.data).toBe(undefined)
  })
})

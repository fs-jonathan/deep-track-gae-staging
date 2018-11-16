import { shallowMount, createLocalVue } from '@vue/test-utils'
import VueRouter from 'vue-router'
import ShowRecord from '@/components/ShowRecord.vue'
import VueElementLoading from 'vue-element-loading'

const localVue = createLocalVue()
localVue.use(VueRouter)
localVue.component('VueElementLoading', VueElementLoading)

const router = new VueRouter()

describe('ShowRecord.vue', () => {
  it('renders correctly', () => {
    const wrapper = shallowMount(ShowRecord, {
      localVue,
      router
    })
    expect(wrapper.isVueInstance()).toBeTruthy()
  })
})

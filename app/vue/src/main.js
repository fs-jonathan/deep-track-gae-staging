import Vue from 'vue'
import App from './App.vue'

// BootstrapVueの読み込み
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

import 'tailwindcss/dist/tailwind.min.css'

import VCalendar from 'v-calendar'
import 'v-calendar/lib/v-calendar.min.css'

import VueMoment from 'vue-moment'
import VueRouter from 'vue-router'

import VueElementLoading from 'vue-element-loading'

Vue.config.productionTip = false

// BootstrapVueの使用
Vue.use(BootstrapVue)
Vue.use(VCalendar, {
  locale: "ja",
  paneWidth: 300
})
Vue.use(VueMoment)
Vue.use(VueRouter)

Vue.component('VueElementLoading', VueElementLoading)

new Vue({
  render: h => h(App),
}).$mount('#app')

import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '@/components/Login.vue'
import GetRecord from '@/components/GetRecord.vue'
import ShowRecord from '@/components/ShowRecord.vue'

Vue.use(VueRouter);

const router = new VueRouter({
  routes: [
    { path: '/login', component: Login },
    { path: '/getRecord', component: GetRecord },
    { path: '/showRecord/:index', name: 'ShowRecord', component: ShowRecord }
  ]
});

export default router;

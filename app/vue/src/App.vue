<template>
  <div id="app">
    <h3>{{ errorMessage }}</h3>
    <vue-element-loading :active="loading" spinner="bar-fade-scale" :is-full-screen="true"/>
    <!-- route outlet -->
    <!-- component matched by the route will render here -->
    <router-view></router-view>
  </div>
</template>

<script>
import VueRouter from 'vue-router'
import axios from 'axios'
import $ from 'liff'

import Login from './components/Login.vue'
import GetRecord from './components/GetRecord.vue'
import ShowRecord from './components/ShowRecord.vue'

const router = new VueRouter({
  routes: [
    { path: '/login', component: Login },
    { path: '/getRecord', component: GetRecord, props: { index: 0 } },
    { path: '/showRecord/:index', name: 'ShowRecord', component: ShowRecord }
  ]
})

export default {
  router,
  name: 'app',
  data: () => ({
    loading: false,
    errorMessage: ""
  }),
  mounted() {
    this.initLiff();
    // this.initData("Hello");
    // this.showLogin();
    // this.showData();
  },
  methods: {
    initLiff: function() {
      var vm = this; // keep reference of viewmodel object
      $.init(function (data) {
        vm.initData(data);
      }, function() {
        vm.errorMessage = "Liff Error";
      });
    },
    initData: function(data) {
      if (data) {
        const userId = data.context.userId;
        this.loading = true

        axios.post('/loginLiff', { "lineUserId": userId })
             .then(() => {
               this.showData();
             })
             .catch(() => {
               this.showLogin();
             })
             .finally(() => {
               this.loading = false
             })
      } else {
        this.errorMessage = "Wrong Page Access";
      }
    },
    showLogin: function() {
      router.replace('/login');
    },
    showData: function() {
      router.replace('/getRecord');
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: left;
  color: #2c3e50;
  margin-top: 60px;
}
</style>

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
import $ from 'liff'

export default {
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

        this.$axios.post('/loginLiff', { "lineUserId": userId })
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
      this.$router.replace('/login');
    },
    showData: function() {
      this.$router.replace('/getRecord');
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

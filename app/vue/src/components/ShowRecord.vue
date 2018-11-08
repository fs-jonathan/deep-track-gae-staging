<template>
  <div class="showRecord">
    <div id="loader">
      <vue-element-loading :active="loading" spinner="bar-fade-scale" :is-full-screen="true"/>
    </div>

    <div class="bg-blue-lightest border-t border-b border-blue text-blue-dark px-4 py-3" v-if="error">
      <p class="text-sm">{{ error }}</p>
    </div>

    <div style="display:none" id="content" v-if="results">
      <div v-for="(result, key, index) in results" :key="index">
        <div class="bg-white hover:bg-grey-light m-3 max-w-sm shadow-md rounded-md overflow-hidden">
          <div class="text-left p-1 sm:text-left sm:flex-grow">
            <span class="text-sm leading-tight">{{ result.title }}</span>
            <span class="text-xs leading-tight text-grey-dark sm:inline">{{ result.subtitle }}</span>
          </div>
          <div class="p-1">
            <strong class="text-xl leading-tight">¥{{ result.cost }}</strong>
          </div>
          <div class="text-left sm:text-left sm:flex-grow p-1">
            <img src="../assets/icon_arrow_right.png" class="w-2"/>
            <span class="text-xs text-grey-dark m-1">¥{{ result.compare }}({{ result.rate }}%)</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ShowRecord',
  data: () => ({
    loading: false,
    results: null,
    error: null
  }),
  created () {
    // fetch the data when the view is created and the data is already being observed
    this.getJson(this.$route.params.index)
  },
  methods: {
    getJson: function(id) {
      this.loading = true

      axios.post('/getJson', { "message": id })
           .then(response => {
             this.results = response.data;
           })
           .catch(err => {
             this.error = err.toString();
           })
           .finally(() => {
             this.loading = false
             document.getElementById("content").style.display = "block";
           })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.arrow {
  width: 12px;
  height: 12px;
}
</style>

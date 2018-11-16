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
        <DetailRow :result="result" />
      </div>
    </div>
  </div>
</template>

<script>
import DetailRow from './DetailRow'

export default {
  name: 'ShowRecord',
  data: () => ({
    loading: false,
    results: null,
    error: null
  }),
  components: {
    DetailRow,
  },
  created () {
    // fetch the data when the view is created and the data is already being observed
    this.getJson(this.$route.params.index)
  },
  methods: {
    getJson: function(id) {
      this.loading = true

      this.$axios.post('/getDetail', { "message": id })
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

<style scoped>
</style>

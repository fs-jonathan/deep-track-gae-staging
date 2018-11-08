<template>
  <div class="getRecord">
    <div id="loader">
      <vue-element-loading :active="loading" spinner="bar-fade-scale" :is-full-screen="true"/>
    </div>

    <div class="bg-blue-lightest border-t border-b border-blue text-blue-dark px-4 py-3" v-if="error">
      <p class="text-sm">{{ error }}</p>
    </div>

    <div id="calendar" class='control border-b m-3'>
      <v-date-picker
        mode='range'
        v-model='selectedDate'
        :input-props='{ class: "input flex", readonly: true, style: "min-width: 300px;" }'
        show-caps>
        <b-field>
          <b-input
            type='text'
            icon='calendar'
            :value='inputValue'
            rounded>
          </b-input>
        </b-field>
      </v-date-picker>
    </div>

    <div style="display:none" id="content" v-if="results">
      <div v-for="(result, key, index) in results" :key="index">
        <div class="bg-white hover:bg-grey-light m-3 max-w-sm shadow-md rounded-md overflow-hidden" v-on:click="showRecord(result.id)">
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
import axios from 'axios'
import moment from 'moment'

export default {
  name: 'GetRecord',
  data: () => ({
    loading: false,
    results: null,
    error: null,
    selectedDate: {
      start: moment(Date.now()).subtract(5, 'd').toDate(),
      end: new Date()
    }
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
    },
    showRecord: function(id) {
      this.$router.push({ name: 'ShowRecord', params: { index: id } });
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

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
        <row :result="result" />
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import moment from 'moment'
import BaseRow from './BaseRow'

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
  components: {
    'row': BaseRow,
  },
  created () {
    // fetch the data when the view is created and the data is already being observed
    this.getJson()
  },
  methods: {
    getJson: function() {
      this.loading = true

      axios.post('/getReport')
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

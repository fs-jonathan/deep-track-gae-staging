<template>
  <div class="bg-white hover:bg-grey-light m-3 max-w-sm shadow-md rounded-md overflow-hidden" v-on:click="showRecord(result.id)">
    <div class="text-left p-1 sm:text-left sm:flex-grow">
      <span class="text-sm leading-tight">{{ result.title }}</span>
      <span class="text-xs leading-tight text-grey-dark sm:inline">{{ result.subtitle }}</span>
    </div>
    <div class="p-1">
      <strong class="text-xl leading-tight">¥{{ result.cost }}</strong>
    </div>
    <div class="text-left sm:text-left sm:flex-grow p-1" v-if="result.compare > 0">
      <img src="../assets/icon_arrow_right.png" class="w-2"/>
      <span class="text-xs text-grey-dark m-1">¥{{ result.compare }}({{ result.rate }}%)</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BaseRow',
  props: ['result'],
  methods: {
    showRecord: function(id) {
      this.$router.push({ name: 'ShowRecord', params: { index: id } });
    }
  }
}
</script>

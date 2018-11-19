<template>
  <div class="bg-white hover:bg-grey-light m-3 max-w-sm shadow-md rounded-md overflow-hidden" v-on:click="showRecord(result.id)">
    <div class="text-left p-1 sm:text-left sm:flex-grow">
      <span class="text-sm leading-tight">{{ this.getTitle(result.id) }}</span>
      <span class="text-xs leading-tight text-grey-dark sm:inline">{{ this.getSubtitle(result.id) }}</span>
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
    },
    getSubtitle: function(id) {
      switch (id) {
        case 1: return "（現時点まで）";
        case 2: return "先週と同じ曜日の比較";
        case 3: return "先月と同じ日時の比較";
        default: return "";
      }
    },
    getTitle: function(id) {
      switch (id) {
        case 1: return "本日";
        case 2: return "昨日";
        case 3: return "今月（現時点まで）";
        case 4: return "先月";
        case 5: return "全期間";
        default: return "";
      }
    }
  }
}
</script>

<style scoped>
</style>

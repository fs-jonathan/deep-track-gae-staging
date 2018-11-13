<template>
  <div class="bg-white hover:bg-grey-light m-3 max-w-sm shadow-md rounded-md overflow-hidden">
    <div class="text-left p-1 sm:text-left sm:flex-grow">
      <span class="text-sm leading-tight">{{ result.title }}</span>
      <span class="text-xl leading-tight" style="float:right" v-if="result.id <= 8">¥{{ result.cost.toFixed(0) }}</span>
      <span class="text-xl leading-tight" style="float:right" v-else>{{ result.cost.toFixed(2) }}%</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DetailRow',
  props: ['result']
}
</script>

<template>

  <a-form-item :label="$t('limit')" help="default 5">
    <a-input-number class="numberInput" v-model="data.limit" :min="0" :max="100" />
  </a-form-item>

  <a-form-item label="query size" help="default 100">
    <a-input-number class="numberInput" v-model="data.query_size" :min="1"  />
  </a-form-item>

  <a-form-item label="min width" help="default 240">
    <a-input-number class="numberInput" v-model="data.min_width" :min="1"  />
  </a-form-item>

  <a-form-item label="min height" help="default 150">
    <a-input-number class="numberInput" v-model="data.min_height" :min="1"  />
  </a-form-item>

  <a-form-item label="max height ratio" help="default 0.7">
    <a-input-number class="numberInput" v-model="data.max_height_ratio" :min="0.1" :precision="1" :step="0.1" />
    <span class="text-gray-400 ml-2" style="font-size:12px">height/width ＜ {{ data.max_height_ratio }}</span>
  </a-form-item>


  <a-form-item label="down timeout" help="default 20">
    <a-input-number class="numberInput" v-model="data.down_timeout" :min="1"  /> <span class="text-gray-400 ml-2" style="font-size:12px">second</span>
  </a-form-item>

</template>


<script setup>
import {inject} from "vue";
const data = inject("options")

</script>

<style scoped>
.numberInput{
  width: 140px;
}
</style>

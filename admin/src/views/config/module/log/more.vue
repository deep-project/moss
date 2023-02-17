<template>
  <div class="p-3">
    <a-form-item label="http pool size"><a-input-number class="w-64" v-model="data.http_pool_size" /></a-form-item>
    <a-form-item label="http pool size"><a-input-number class="w-64" v-model="data.sql_pool_size" /></a-form-item>
  </div>
</template>
<script setup>
import {inject} from "vue";

const data = inject('data')
</script>
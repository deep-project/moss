<template>

  <a-form-item label="referer">
    <a-input v-model="data.referer" class="input" />
  </a-form-item>

  <a-form-item label="proxy">
    <a-input v-model="data.proxy" class="input" />
  </a-form-item>

  <a-form-item label="userAgent">
    <a-textarea v-model="data.user_agent" />
  </a-form-item>

  <a-form-item label="retry">
    <a-input-number v-model="data.retry" class="input" :min="0" />
  </a-form-item>
</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")
</script>
<style scoped>
.input{
  width: 220px;
}
</style>


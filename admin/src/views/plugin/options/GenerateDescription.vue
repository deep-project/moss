<template>

  <a-form-item :label="$t('enable')">
    <a-switch type="round" v-model="data.enable" />
  </a-form-item>

  <a-form-item label="length">
    <a-input-number v-model="data.length" :style="{ width: '200px' }" :min="1" :max="250" />
  </a-form-item>
</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")
</script>

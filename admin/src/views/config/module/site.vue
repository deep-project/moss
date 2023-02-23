<template>
  <a-form-item :label="$t('name')">
    <a-input v-model="data.name" class="w-64" />
  </a-form-item>

  <a-form-item :label="$t('url')">
    <a-input  class="w-64" v-model="data.url" placeholder="http://www.xxx.com" />
  </a-form-item>

  <a-form-item :label="$t('title')" style="max-width: 600px">
    <a-input v-model="data.title" />
  </a-form-item>

  <a-form-item :label="$t('keywords')" field="keywords" style="max-width: 600px" extra="keywords">
    <a-textarea v-model="data.keywords" class="w-full" :auto-size="{ minRows: 3, maxRows: 3 }"/>
  </a-form-item>

  <a-form-item :label="$t('description')" field="description" style="max-width: 600px" extra="description">
    <a-textarea v-model="data.description" class="w-full" :auto-size="{ minRows: 3, maxRows: 3 }"/>
  </a-form-item>

  <a-divider />
</template>

<script setup>
  import {inject,watch} from 'vue'

  const data = inject('data')

  watch(()=>data.value.url, (val)=>{
    if(!val) return
    if(!val.startsWith("http") && !val.startsWith("//")) data.value.url = "http://" + data.value.url
  })

</script>
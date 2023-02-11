<template>
    <a-form-item label="endpoint">
      <a-input class="w-64" v-model="data.endpoint" />
      <span class="ml-2 text-gray-400">s3-us-west-1.amazonaws.com</span>
    </a-form-item>

    <a-form-item label="region">
      <a-input class="w-64" v-model="data.region" />
      <span class="ml-2 text-gray-400">us-west-1</span>
    </a-form-item>

    <a-form-item label="bucket">
      <a-input class="w-64" v-model="data.bucket" />
    </a-form-item>

    <a-form-item label="accessKey">
      <a-input class="w-64" v-model="data.accessKey" />
    </a-form-item>

    <a-form-item label="secretKey">
      <a-input-password class="w-64" v-model="data.secretKey" />
    </a-form-item>
</template>

<script setup>
  defineProps({data:Object})
</script>
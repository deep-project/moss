<template>

    <a-form-item label="endpoint">
      <a-input class="w-64" v-model="data.endpoint" />
    </a-form-item>

    <a-form-item label="bucket">
      <a-input class="w-64" v-model="data.bucket" />
    </a-form-item>

    <a-form-item label="accessKeyID">
      <a-input class="w-64" v-model="data.accessKeyID" />
    </a-form-item>

    <a-form-item label="accessKeySecret">
      <a-input-password class="w-64" v-model="data.accessKeySecret" />
    </a-form-item>

</template>

<script setup>
  defineProps({data:Object})
</script>
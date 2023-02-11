<template>

    <a-form-item label="bucketURL">
      <a-input class="w-64" v-model="data.bucketURL" />
    </a-form-item>

    <a-form-item label="secretID">
      <a-input class="w-64" v-model="data.secretID" />
    </a-form-item>

    <a-form-item label="secretKey">
      <a-input-password class="w-64" v-model="data.secretKey" />
    </a-form-item>

</template>

<script setup>
  defineProps({data:Object})
</script>
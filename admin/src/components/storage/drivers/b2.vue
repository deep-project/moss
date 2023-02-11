<template>

    <a-form-item label="keyID">
      <a-input class="w-64" v-model="data.keyID" />
    </a-form-item>

    <a-form-item label="bucket">
      <a-input class="w-64" v-model="data.bucket" />
    </a-form-item>

    <a-form-item label="appKey">
      <a-input-password class="w-64" v-model="data.appKey" />
    </a-form-item>

</template>

<script setup>

  import {useStore} from "@/store";

  defineProps({data:Object})
  const store = useStore()
</script>
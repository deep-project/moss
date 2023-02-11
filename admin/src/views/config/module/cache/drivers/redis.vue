<template>
    <a-form-item label="addr">
      <a-input class="w-64" v-model="data.addr" />
    </a-form-item>
    <a-form-item label="password">
      <a-input-password class="w-64" v-model="data.password"  />
    </a-form-item>
    <a-form-item label="db">
      <a-input-number v-model="data.db" :min="0" class="w-64" />
    </a-form-item>
</template>

<script setup>
  defineProps({data:Object})
</script>
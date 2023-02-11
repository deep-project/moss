<template>
    <a-form-item label="addr">
      <a-input class="w-64" v-model="data.addr" />
    </a-form-item>
</template>

<script setup>
  defineProps({data:Object})
</script>
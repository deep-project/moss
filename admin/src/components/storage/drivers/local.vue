<template>
    <a-form-item label="path">
      <a-input class="w-64" v-model="data.path" />
    </a-form-item>
</template>

<script setup>
  defineProps({data:Object})
</script>
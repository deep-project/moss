<template>
    <a-form-item label="host">
      <a-input class="w-64" v-model="data.host" />
    </a-form-item>
    <a-form-item label="port">
      <a-input class="w-64" v-model="data.port"/>
    </a-form-item>
    <a-form-item label="name">
      <a-input class="w-64" v-model="data.name" />
    </a-form-item>
    <a-form-item label="password">
      <a-input-password class="w-64" v-model="data.password" />
    </a-form-item>
    <a-form-item label="timeout">
      <a-input-number class="w-64" v-model="data.timeout" /> <span class="text-gray-300 ml-1">second</span>
    </a-form-item>

</template>

<script setup>
  defineProps({data:Object})
</script>
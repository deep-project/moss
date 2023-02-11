<template>
  <a-breadcrumb>
    <template v-if="route.name ==='dashboard'">
      <a-breadcrumb-item :clickable="false">{{ $t('dashboard') }}</a-breadcrumb-item>
      <a-breadcrumb-item>{{ $t('home') }}</a-breadcrumb-item>
    </template>
    <template v-else>
      <a-breadcrumb-item><router-link :to="{name:'index'}">{{ $t('dashboard') }}</router-link></a-breadcrumb-item>
      <a-breadcrumb-item>{{ $t(String(route.name).split("-")[0]) }}</a-breadcrumb-item>
    </template>
  </a-breadcrumb>
</template>

<script setup>

  import {useRoute} from "vue-router";

  const route = useRoute()

</script>
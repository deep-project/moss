<template>
  <a-button @click="store.siderCollapsed = !store.siderCollapsed">
    <template #icon>
      <icon-menu-unfold :size="18" v-if="store.siderCollapsed" />
      <icon-menu-fold :size="18" v-else />
    </template>
  </a-button>
</template>

<script setup>
  import {useStore} from "@/store";

  const store = useStore()
</script>
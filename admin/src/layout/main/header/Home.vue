<template>
  <a-button @click="open">
    <template #icon>
      <icon-home :size="18" />
    </template>
  </a-button>
</template>

<script setup>
  import {useOpenLink} from "@/hooks/utils.js";

  function open(){
    useOpenLink(window.location.origin,'')
  }

</script>
<template>
  <a-button @click="open">
    <template #icon>
      <icon-home :size="18" />
    </template>
  </a-button>
</template>

<script setup>
  import {useOpenLink} from "@/hooks/utils.js";
  import {useSiteURL} from "@/hooks/app/index.js"
  import {useStore} from "@/store/index.js";


  const store = useStore()
  function open(){
    useOpenLink(useSiteURL(store),'')
  }

</script>
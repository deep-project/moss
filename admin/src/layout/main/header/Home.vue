<template>
  <a-button @click="open">
    <template #icon>
      <icon-home :size="18" />
    </template>
  </a-button>
</template>

<script setup>
import {useStore} from "@/store/index.js";
import {useOpenLink} from "@/hooks/utils.js";
import {useSiteURL} from "@/hooks/app/index.js"


const store = useStore()
const open = ()=>{ useOpenLink(useSiteURL(store),'') }

</script>
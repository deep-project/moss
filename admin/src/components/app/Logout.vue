<template>
  <a-popconfirm position="br" :content="$t('message.logoutSure')" @ok="store.token = ''">
    <a-button>
      <template #icon><icon-export :size="18" /></template>
    </a-button>
  </a-popconfirm>

</template>

<script setup>

  import {useStore} from "@/store/index.js";

  const store = useStore()
</script>
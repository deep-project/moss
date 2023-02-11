<template>
  <a-dropdown @select="useSetLocale" position="br">
    <a-button :type="type">
      <template #icon>
        <icon-language :size="18" />
      </template>
    </a-button>
    <template #content>
      <a-doption v-for="(locale ,key) in useLocaleList()" :value="key">{{ locale.name }}</a-doption>
    </template>
  </a-dropdown>
</template>

<script setup>
  import {useSetLocale,useLocaleList} from "@/locale";

  defineProps({type:String})
</script>
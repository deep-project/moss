<template>
  <a-dropdown class="locale" @select="useSetLocale" position="br" :scrollbar="false" style="height:100%">
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


<style>
.locale .arco-dropdown-list-wrapper{
  max-height:500px !important;
}
</style>
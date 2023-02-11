<template>
  <a-card class="w-full" :title="$t('about') + ' Moss'" :bordered="false">
    <div class="mb-2 leading-8">
      Moss 苔藓、青苔之意，希望moss可以如青苔一样，顽强的活着，遍地生根。<br>
      本程序基于MIT协议开源，完全免费，请放心使用。不明白的地方可以与作者沟通。切勿用于非法用途。
    </div>
    <a-link @click="useOpenLink">
      <template #icon><icon-github /></template>https://github.com/deep-project/moss
    </a-link>
  </a-card>
</template>

<script setup>
  import {useOpenLink} from '@/hooks/utils.js'
</script>
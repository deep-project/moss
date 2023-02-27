<template>
  <a-card class="w-full" :title="$t('about') + ' Moss'" :bordered="false">
    <div class="mb-2 leading-8">
      {{ $t('mossAbout') }}<br>{{ $t('mossInstructions') }}
    </div>
    <a-link @click="useOpenLink">
      <template #icon><icon-github /></template>https://github.com/deep-project/moss
    </a-link>
  </a-card>
</template>

<script setup>
  import {useOpenLink} from '@/hooks/utils.js'
</script>
<template>
  <a-card class="w-full" :title="$t('about') + ' Moss'" :bordered="false">
    <div v-if="appInfo" class="mb-2 leading-8 italic">
      version {{ appInfo?.version }}
    </div>
    <div class="mb-2 leading-8">
      {{ $t('mossAbout') }}<br>{{ $t('mossInstructions') }}
    </div>
    <a-link @click="useOpenLink">
      <template #icon><icon-github /></template>https://github.com/deep-project/moss
    </a-link>
  </a-card>
</template>

<script setup>
  import {useRequest} from "vue-request";
  import {useOpenLink} from '@/hooks/utils.js'
  import {dashboardData} from "@/api/index.js";


  const { data:appInfo } = useRequest(dashboardData, {defaultParams:['appInfo']})
</script>
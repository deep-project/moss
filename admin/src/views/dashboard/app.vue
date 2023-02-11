<template>
  <a-card class="w-full h-full" :title="$t('app')" :bordered="false" :body-style="{height:'calc(100% - 46px)'}">

    <a-grid :cols="{ xs: 2, sm: 3 }" class="h-full items-center">

      <a-grid-item>
        <a-statistic :title="$t('database')" :value="database[0]" :value-from="0" :precision="isInt(database[0]) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">{{ database[1] }}</b></template>
          <template #extra><a-spin v-if="loadingDatabase" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item>
        <a-statistic :title="$t('log')" :value="log[0]" :value-from="0" :precision="isInt(log[0]) ? 0:2" show-group-separator animation>
          <template #suffix><span class="ml-1">{{ log[1] }}</span></template>
          <template #extra><a-spin v-if="loadingLog" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item v-if="cache[0] >= 0">
        <a-statistic :title="$t('cache')" :value="cache[0]" :value-from="0" :precision="isInt(cache[0]) ? 0:2" show-group-separator animation>
          <template #suffix><span class="ml-1">{{ cache[1] }}</span></template>
          <template #extra><a-spin v-if="loadingCache" /></template>
        </a-statistic>
      </a-grid-item>

    </a-grid>

  </a-card>
</template>



<script setup>
  import {useRequest} from "vue-request";
  import {dashboard} from "@/api/index.js";
  import {computed} from "vue";
  import {useParseBytesSize} from "@/hooks/utils.js";

  const { data:databaseBytes, loading:loadingDatabase } = useRequest(dashboard, {defaultParams:['database']})
  const database = computed(()=>useParseBytesSize(databaseBytes.value))

  const { data:logBytes, loading:loadingLog } = useRequest(dashboard, {defaultParams:['log']})
  const log = computed(()=>useParseBytesSize(logBytes.value))

  const { data:cacheBytes, loading:loadingCache } = useRequest(dashboard, {defaultParams:['cache']})
  const cache = computed(()=> useParseBytesSize(cacheBytes.value))


  function isInt(val){
    if(!val) return false
    return parseInt(val) === parseFloat(val.toFixed(2))
  }
</script>
<template>
  <a-card class="w-full h-full" :title="$t('app')" :bordered="false" :body-style="{height:'calc(100% - 46px)'}">
    <a-grid :cols="{ xs: 2, sm: 3, md: 4, lg:5 }" :colGap="12" :rowGap="12">
      <a-grid-item class="flex items-center" style="min-height: 104px">
        <a-statistic title="cpu" :value="cpu" :value-from="oldCPU" :precision="isInt(cpu) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">%</b></template>
          <template #extra></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item class="flex items-center" style="min-height: 104px">
        <a-statistic :title="$t('memory')" :value="memory[0]" :value-from="oldMemory" :precision="isInt(memory[0]) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">{{ memory[1] }}</b></template>
          <template #extra></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item class="flex items-center" style="min-height: 104px">
        <a-statistic :title="$t('database')" :value="database[0]" :value-from="0" :precision="isInt(database[0]) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">{{ database[1] }}</b></template>
          <template #extra><a-spin v-if="loadingDatabase" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item class="flex items-center" style="min-height: 104px">
        <a-statistic :title="$t('log')" :value="log[0]" :value-from="0" :precision="isInt(log[0]) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">{{ log[1] }}</b></template>
          <template #extra><a-spin v-if="loadingLog" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item v-if="cache[0] >= 0" class="flex items-center" style="min-height: 104px">
        <a-statistic :title="$t('cache')" :value="cache[0]" :value-from="0" :precision="isInt(cache[0]) ? 0:2" show-group-separator animation>
          <template #suffix><b class="ml-1">{{ cache[1] }}</b></template>
          <template #extra><a-spin v-if="loadingCache" /></template>
        </a-statistic>
      </a-grid-item>

    </a-grid>

  </a-card>
</template>



<script setup>
  import {useRequest} from "vue-request";
  import {dashboardData} from "@/api/index.js";
  import {computed,ref,watch} from "vue";
  import {useParseBytesSize} from "@/hooks/utils.js";


  const oldCPU = ref(0)
  const { data:cpu } = useRequest(dashboardData, {defaultParams:['appCPU'], pollingInterval: 2000, errorRetryCount: 1 })
  watch(cpu, (n,old)=>{ oldCPU.value = old })

  const oldMemory = ref(0)
  const { data:memoryBytes } = useRequest(dashboardData, {defaultParams:['appMemory'], pollingInterval: 2000, errorRetryCount: 1 })
  const memory = computed(()=>useParseBytesSize(memoryBytes.value))
  watch(memory,(n,old)=>{ oldMemory.value = old[0] })

  const { data:databaseBytes, loading:loadingDatabase } = useRequest(dashboardData, {defaultParams:['database']})
  const database = computed(()=>useParseBytesSize(databaseBytes.value))

  const { data:logBytes, loading:loadingLog } = useRequest(dashboardData, {defaultParams:['log']})
  const log = computed(()=>useParseBytesSize(logBytes.value))

  const { data:cacheBytes, loading:loadingCache } = useRequest(dashboardData, {defaultParams:['cache']})
  const cache = computed(()=> useParseBytesSize(cacheBytes.value))


  function isInt(val){
    if(!val) return false
    return parseInt(val) === parseFloat(val.toFixed(2))
  }
</script>
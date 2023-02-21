<template>
  <a-card class="w-full" :title="$t('storehouse')" :bordered="false">

    <a-grid :cols="{ xs: 2, sm: 3, md: 4 }" :colGap="12" :rowGap="12">
      <a-grid-item>
        <a-statistic :title="$t('total')" :value="total" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('today')" :value="today" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingToday" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('yesterday')" :value="yesterday" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingYesterday" /></template>

        </a-statistic>
      </a-grid-item>
    </a-grid>

  </a-card>
</template>


<script setup>

  import {useRequest} from "vue-request";
  import {dashboardData} from "@/api/index.js";

  const { data:total, loading:loadingTotal } = useRequest(dashboardData, {defaultParams:['storeTotal']})
  const { data:today, loading:loadingToday } = useRequest(dashboardData, {defaultParams:['storeToday']})
  const { data:yesterday, loading:loadingYesterday } = useRequest(dashboardData, {defaultParams:['storeYesterday']})


</script>
<template>
  <a-card class="w-full" :title="$t('article')" :bordered="false">

    <a-grid :cols="{ xs: 2, sm: 3, md: 4, lg:5 }" :colGap="12" :rowGap="12">

      <a-grid-item>
        <a-statistic :title="$t('today')" :value="today" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingToday" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item>
        <a-statistic :title="$t('yesterday')" :value="yesterday" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingYesterday" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item>
        <a-statistic :title="'7 '+$t('days')" :value="days7" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingDays7" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item>
        <a-statistic :title="'30 '+$t('days')" :value="days30" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-file /></span></template>
          <template #extra><a-spin v-if="loadingDays30" /></template>
        </a-statistic>
      </a-grid-item>

      <a-grid-item>
        <a-statistic :title="$t('total')" :value="total" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-file class="" /></span></template>
          <template #extra><a-spin v-if="loadingTotal" /></template>
        </a-statistic>
      </a-grid-item>


    </a-grid>

  </a-card>
</template>


<script setup>

  import {useRequest} from "vue-request";
  import {dashboardData} from "@/api/index.js";

  const { data:total, loading:loadingTotal } = useRequest(dashboardData, {defaultParams:['articleTotal']})
  const { data:today, loading:loadingToday } = useRequest(dashboardData, {defaultParams:['articleToday']})
  const { data:yesterday, loading:loadingYesterday } = useRequest(dashboardData, {defaultParams:['articleYesterday']})
  const { data:days7, loading:loadingDays7 } = useRequest(dashboardData, {defaultParams:['articleLast7days']})
  const { data:days30, loading:loadingDays30 } = useRequest(dashboardData, {defaultParams:['articleLast30days']})

</script>
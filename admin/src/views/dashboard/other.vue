<template>
  <a-card class="w-full" :title="$t('other')" :bordered="false">

    <a-grid :cols="{ xs: 2, sm: 3, md: 4, lg:5 }" :colGap="12" :rowGap="12">
      <a-grid-item>
        <a-statistic :title="$t('storehouse')" :value="storeTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-storage /></span></template>
          <template #extra><a-spin v-if="loadingStoreTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('category')" :value="categoryTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-list /></span></template>
          <template #extra><a-spin v-if="loadingCategoryTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('tag')" :value="tagTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-tag /></span></template>
          <template #extra><a-spin v-if="loadingTagTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('link')" :value="linkTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1 text-gray-500"><icon-link /></span></template>
          <template #extra><a-spin v-if="loadingLinkTotal" /></template>

        </a-statistic>
      </a-grid-item>
    </a-grid>

  </a-card>
</template>


<script setup>

  import {useRequest} from "vue-request";
  import {dashboardData} from "@/api/index.js";

  const { data:storeTotal, loading:loadingStoreTotal } = useRequest(dashboardData, {defaultParams:['storeTotal']})
  const { data:categoryTotal, loading:loadingCategoryTotal } = useRequest(dashboardData, {defaultParams:['categoryTotal']})
  const { data:tagTotal, loading:loadingTagTotal } = useRequest(dashboardData, {defaultParams:['tagTotal']})
  const { data:linkTotal, loading:loadingLinkTotal } = useRequest(dashboardData, {defaultParams:['linkTotal']})


</script>
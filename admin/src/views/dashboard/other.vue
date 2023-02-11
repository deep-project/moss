<template>
  <a-card class="w-full" :title="$t('other')" :bordered="false">

    <a-grid :cols="{ xs: 2, sm: 3, md: 4 }" :colGap="12" :rowGap="12">
      <a-grid-item>
        <a-statistic :title="$t('category')" :value="categoryTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-list /></span></template>
          <template #extra><a-spin v-if="loadingCategoryTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('tag')" :value="tagTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-tag /></span></template>
          <template #extra><a-spin v-if="loadingTagTotal" /></template>
        </a-statistic>
      </a-grid-item>
      <a-grid-item>
        <a-statistic :title="$t('link')" :value="linkTotal" :value-from="0" show-group-separator animation>
          <template #suffix><span class="ml-1"><icon-link /></span></template>
          <template #extra><a-spin v-if="loadingLinkTotal" /></template>

        </a-statistic>
      </a-grid-item>
    </a-grid>

  </a-card>
</template>


<script setup>

  import {useRequest} from "vue-request";
  import {dashboard} from "@/api/index.js";

  const { data:categoryTotal, loading:loadingCategoryTotal } = useRequest(dashboard, {defaultParams:['categoryTotal']})
  const { data:tagTotal, loading:loadingTagTotal } = useRequest(dashboard, {defaultParams:['tagTotal']})
  const { data:linkTotal, loading:loadingLinkTotal } = useRequest(dashboard, {defaultParams:['linkTotal']})


</script>
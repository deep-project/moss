<template>

  <div class="w-full" style="max-width: 360px">

    <a-form-item :label="$t('limit')">
      <a-input-number v-model="data.limit" :min="0"></a-input-number>
    </a-form-item>

    <a-form-item :label="$t('order')" v-if="!disableOrder">
      <a-select v-model="data.order">
        <a-option value="">none</a-option>
        <a-option value="id desc">id desc</a-option>
        <a-option value="id asc">id asc</a-option>
        <a-option value="views desc">views desc</a-option>
        <a-option value="views asc">views asc</a-option>
      </a-select>
    </a-form-item>

    <template v-if="!showPageOption">
      <a-form-item :label="$t('range')">
        <a-tree-select
            v-model="data.category_ids"
            :allow-search="true"
            :allow-clear="true"
            :tree-checkable="true"
            :tree-check-strictly="true"
            :loading="loading"
            :data="treeData"
            :fieldNames="{key: 'id',title: 'name'}"
        ></a-tree-select>
      </a-form-item>
    </template>
    <template v-else>
      <a-form-item :label="$t('maxPage')">
        <a-input-number v-model="data.max_page" :min="0"></a-input-number>
      </a-form-item>
      <a-form-item label="disable count">
        <a-switch type="round" v-model="data.disable_count" />
      </a-form-item>
    </template>

  </div>

</template>

<script setup>
  import {useRequest} from "vue-request";
  import {categoryTree} from "@/api/index.js";

  const { data,showPageOption,disableOrder } = defineProps({data:Object, showPageOption:Boolean, disableOrder:Boolean})
  const {data:treeData,loading} = useRequest(categoryTree);

</script>
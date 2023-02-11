<template>
  <a-form-item :label="$t('select')">
    <a-tree-select
        v-model="data.menu.select"
        :allow-search="true"
        :allow-clear="true"
        :tree-checkable="true"
        :tree-check-strictly="true"
        :loading="loading"
        :data="treeData"
        :fieldNames="{key: 'id',title: 'name'}"
        style="max-width: 350px"
    ></a-tree-select>
  </a-form-item>

</template>

<script setup>
  import {inject} from 'vue'
  import {useRequest} from "vue-request";
  import {categoryTree} from "@/api/index.js";

  const data = inject('data')
  const {data:treeData,loading} = useRequest(categoryTree);

</script>
<template>


  <a-form-item :label="$t('select')">
    <a-space direction="vertical">
      <a-tree-select
          v-model="data.menu.select"
          allow-search allow-clear allow-create scrollbar
          tree-checkable tree-check-strictly
          :loading="loading"
          :data="treeData"
          :fieldNames="{key: 'id',title: 'name'}"
          :style="{width:'320px'}"
          @press-enter="inputPressEnter"
          v-model:input-value="inputValue"
      >
      </a-tree-select>
    </a-space>
  </a-form-item>

  <a-form-item :label="$t('limit')">
    <a-input-number v-model="data.menu.limit" :min="0" style="width: 150px"></a-input-number>
  </a-form-item>

</template>

<script setup>
import {inject,ref} from 'vue'
import {useRequest} from "vue-request";
import {categoryTree} from "@/api/index.js";

const data = inject('data')
const {data:treeData,loading} = useRequest(categoryTree);

const inputValue = ref('')
function inputPressEnter(val){
  let n = Number(val)
  inputValue.value = ""
  if(n && n > 0) data.value.menu.select.push(n)
}


</script>


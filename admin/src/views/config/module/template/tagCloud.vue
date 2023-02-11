<template>
  <div class="w-full" style="max-width: 360px">

    <a-form-item :label="$t('limit')">
      <a-input-number v-model="data.tag_cloud.limit" :min="0"></a-input-number>
    </a-form-item>

    <a-form-item :label="$t('order')">
      <a-select v-model="data.tag_cloud.order">
        <a-option value="">none</a-option>
        <a-option value="id desc">id desc</a-option>
        <a-option value="id asc">id asc</a-option>
      </a-select>
    </a-form-item>

    <a-form-item :label="$t('select')">
      <a-select v-model="data.tag_cloud.select" multiple
                :options="tagSelectList"
                :field-names="{value: 'id', label: 'name'}"
                @search="getTagSelectList"
                :filter-option="false"
                :loading="loadingGetSelected"
                :fallback-option="fallback"
                @focus="focus"
      />
    </a-form-item>

  </div>
</template>

<script setup>
  import {inject,ref} from "vue";
  import {useRequest} from "vue-request";
  import {tableList, tagGetByIds} from "@/api/index.js";

  const data = inject('data')

  // 已选中的标签数据
  const { data:selectedData, run:runGetSelected, loading:loadingGetSelected } = useRequest(tagGetByIds, {manual:true})
  function getSelected(){
    if(data.value.tag_cloud.select && data.value.tag_cloud.select.length > 0) runGetSelected(data.value.tag_cloud.select)
  }
  getSelected()


  // 推荐标签列表
  const { data:tagSelectList, run:runTagSelectList, loading:loadingTagSelectList } = useRequest(tableList, {manual:true})

  function getTagSelectList(searchValue){
    let params = {page:1, limit:100, order:""}
    if(searchValue) params.where = {field:'name', operator:'like', value:searchValue}
    runTagSelectList("tag", params)
  }

  const isLoadSuccess = ref(false)
  function focus(){
    if(!isLoadSuccess.value){
      getTagSelectList()
      isLoadSuccess.value = true
    }
  }

  // 回退函数，当默认列表中没有选中的值数据，则使用回退函数赋值
  function fallback(id){
    if(!selectedData.value || selectedData.value.length === 0) return
    for(let i in selectedData.value){
      if(selectedData.value[i].id === id) return selectedData.value[i]
    }
  }

</script>
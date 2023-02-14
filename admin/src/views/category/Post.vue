<template>
  <a-form-item field="name" :label="$t('name')" :rules="[{ required:true, message: $t('message.required',[$t('name')]) }]">
    <a-input v-model="record.name" :max-length="150" allow-clear show-word-limit />
  </a-form-item>

  <a-form-item field="slug" :label="$t('slug')" :rules="[{required:!!record.id, message:$t('message.required',[$t('slug')])}]">
    <div class="w-full">
      <a-input v-model="record.slug" :max-length="150" allow-clear show-word-limit />
      <div v-if="record.slug" class="break-all text-gray-600" style="margin-top:3px;font-size:12px;">
        <div v-if="record.id > 0" @click="useOpenLink" class="cursor-pointer hover:underline underline-offset-4 hover:text-blue-500">{{ slugURL }}</div>
        <div v-else>{{ slugURL }}</div>
      </div>
    </div>
  </a-form-item>

  <a-form-item field="parent_id" :label="$t('parentCategory')">
    <a-input-group class="w-full">
      <a-input-number v-model="record.parent_id" hide-button allow-clear style="width: 130px;" placeholder="id" />
      <a-cascader :options="treeData" v-model="record.parent_id" :loading="loadingTreeData" :field-names="{value: 'id', label: 'name'}" check-strictly :placeholder="$t('select')" />
    </a-input-group>
  </a-form-item>

  <a-form-item field="title" :label="$t('title')">
    <a-input v-model="record.title" :max-length="250" allow-clear show-word-limit />
  </a-form-item>

  <a-form-item field="description" :label="$t('description')">
    <a-textarea class="input" v-model="record.description" :max-length="250" :auto-size="{minRows:3,maxRows:5}" show-word-limit />
  </a-form-item>

  <a-form-item field="keywords" :label="$t('keywords')">
    <a-textarea class="input" v-model="record.keywords" :max-length="250" :auto-size="{minRows:3,maxRows:5}" show-word-limit />
  </a-form-item>

  <a-form-item field="create_time" :label="$t('createTime')">
    <a-date-picker class="w-full" v-model="createTime" value-format="timestamp" show-time @change="(val)=>record.create_time =parseInt(val / 1000)" />
  </a-form-item>

</template>


<script setup>
  import {computed, inject} from "vue";
  import {useRequest} from "vue-request";
  import {categoryTree} from "@/api/index.js";
  import {useStore} from "@/store/index.js";
  import {useOpenLink} from '@/hooks/utils.js'

  const record = inject('record')
  const createTime = computed(()=>record.value.create_time*1000)
  const {data:treeData,loading:loadingTreeData} = useRequest(categoryTree,{onSuccess:()=>{treeDataDisabled(treeData.value)}});
  const store = useStore()

  const slugURL = computed(()=>{
    let path =  store.config.router.category_rule.replace('{slug}', record.value.slug)
    if(path.indexOf('/')!==0) path = "/" + path
    return store.config.site.url + path
  })


  // treeData禁用自身类目和其子类目
  function treeDataDisabled(list){
    for(let i in list){
      if(list[i].id === record.value.id){
        list[i].disabled = true
        //treeDataDisabledChildren(list[i].children)
      }else{
        treeDataDisabled(list[i].children)
      }
    }
  }

  // 禁用全部子分类
  // 在属性选择器情况下，子分类依然会显示，使用此方法禁用子分类
  // function treeDataDisabledChildren(list){
  //   if(list){
  //     for(let i in list){
  //       list[i].disabled = true
  //       treeDataDisabledChildren(list[i].children)
  //     }
  //   }
  // }

</script>


<style scoped>
:deep(.arco-select-view-value){
  font-size:12px !important;;
}
</style>
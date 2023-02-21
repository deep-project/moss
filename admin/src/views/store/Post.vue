<template>
  <a-form-item field="title" :label="$t('title')" :rules="[{ required:true, message: $t('message.required',[$t('title')]) }]">
    <a-input v-model="record.title" :max-length="150" allow-clear show-word-limit />
  </a-form-item>

  <a-form-item field="content" :label="$t('content')" :rules="[{ required:true, message: $t('message.required',[$t('content')]) }]">
    <a-textarea v-model="record.content" :auto-size="{ minRows:7, maxRows:7 }" />
  </a-form-item>

  <a-form-item field="category_id" :label="$t('category')+'('+$t('id')+')'">
    <a-input-number v-model="record.category_id" :min="0" />
  </a-form-item>

  <a-form-item field="category_name" :label="$t('category')+'('+$t('name')+')'">
    <a-input :disabled="record.category_id > 0" v-model="record.category_name" :max-length="250" allow-clear />
  </a-form-item>

  <a-form-item field="tags" :label="$t('tags')">
    <a-input-tag v-model="record.tags" allow-clear placeholder="Please Enter"/>
  </a-form-item>

  <a-form-item field="slug" :label="$t('slug')">
    <div class="w-full">
      <a-input v-model="record.slug" :max-length="250" allow-clear />
      <div v-if="record.slug" class="break-all text-gray-600" style="margin-top:10px;font-size:12px;">{{ slugURL }}</div>
    </div>
  </a-form-item>

  <a-form-item field="thumbnail" :label="$t('thumbnail')">
    <UploadImgInput class="w-full" v-model="record.thumbnail" />
  </a-form-item>

  <a-form-item field="description" :label="$t('description')">
    <a-textarea v-model="record.description" />
  </a-form-item>

  <a-form-item field="keywords" :label="$t('keywords')">
    <a-textarea v-model="record.keywords" />
  </a-form-item>

  <a-form-item field="views" :label="$t('views')">
    <a-input-number v-model="record.views" :min="0" />
  </a-form-item>

  <a-form-item v-if="record.id > 0" field="create_time" :label="$t('createTime')">
    <a-date-picker class="w-full" v-model="createTime" value-format="timestamp" show-time @change="(val)=>record.create_time =parseInt(val / 1000)" />
  </a-form-item>

  <a-collapse expand-icon-position="right">
    <a-collapse-item :header="$t('extends')" key="extends" style="background: transparent">
      <a-form :model="record" :label-col-props="{span: 8}" :wrapper-col-props="{span: 16}">
        <a-form-item v-for="(item,index) in record.extends" :label-col-style="{paddingRight:'10px'}">
          <template #label>
            <div class="flex">
              <a-input class="input input_extends" v-model="item.key" /><span class="ml-2">:</span>
            </div>
          </template>
          <a-textarea class="input input_extends" :auto-size="{minRows:1,maxRows:5}" v-model="item.value" />
          <a-button class="ml-1" type="text" @click="record.extends.splice(index,1)"><template #icon><icon-close-circle :stroke-width="3" /></template></a-button>
        </a-form-item>
      </a-form>
      <a-button size="mini" long @click="addExtends">
        <template #icon><icon-plus /></template>{{$t('add')}}
      </a-button>
    </a-collapse-item>
  </a-collapse>

</template>


<script setup>
  import {computed, inject} from "vue";
  import UploadImgInput from '@/components/utils/UploadImgInput.vue'
  import {useAppendSiteURL} from "@/hooks/app/index.js";
  import {useStore} from "@/store/index.js";

  const record = inject('record')
  const createTime = computed(()=>record.value.create_time*1000)
  const store = useStore()
  if(! record.value.views) record.value.views = 0
  const slugURL = computed(()=> useAppendSiteURL(store, store.config.router.article_rule.replace('{slug}', record.value.slug)))


  function addExtends(){
    if(!record.value.extends) record.value.extends = []
    record.value.extends.push({key:'',value:''})
  }

</script>
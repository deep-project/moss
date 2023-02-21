<template>

    <a-form-item field="slug" :label="$t('slug')" :style="{marginBottom:record.slug ?'10px':''}"
                 :rules="[{required:!!record.id, message:$t('message.required',[$t('slug')])}]" hide-asterisk>
      <div class="w-full">
        <a-input class="input" v-model="record.slug" :max-length="150" allow-clear show-word-limit />
        <div v-if="record.slug" class="break-all text-gray-600" style="margin-top:10px;font-size:12px;">
          <div v-if="record.id > 0" @click="useOpenLink" class="cursor-pointer hover:underline underline-offset-4 hover:text-blue-500">{{ slugURL }}</div>
          <div v-else>{{ slugURL }}</div>
        </div>
      </div>
    </a-form-item>

    <a-form-item field="thumbnail" :label="$t('thumbnail')">
      <div class="w-full" >
        <UploadImgInput v-model="record.thumbnail" class="w-full" inputStyle="background-color: var(--color-bg-5);" />
        <a-card v-if="record.thumbnail" class="w-full mt-5" size="mini" :bordered="false">
          <template #title><span class="text-sm">{{$t('preview')}}</span></template>
          <template #extra>
            <icon-delete class="cursor-pointer" @click="record.thumbnail=''" />
          </template>
          <div class="text-center"><a-image :src="record.thumbnail" height="170" width="100%" referrerpolicy="no-referrer" /></div>
        </a-card>
      </div>
    </a-form-item>

    <a-form-item field="category_id" :label="$t('category')">
      <SelectCategory v-model="record.category_id" :cascader-style="{backgroundColor:'var(--color-bg-5)'}" />
    </a-form-item>

    <a-form-item :label="$t('tag')">
      <Tag />
    </a-form-item>

    <a-form-item field="description" :label="$t('description')">
      <a-textarea class="input" v-model="record.description" :max-length="250" :auto-size="{minRows:3,maxRows:5}" show-word-limit />
    </a-form-item>

    <a-form-item field="keywords" :label="$t('keywords')">
      <a-textarea class="input" v-model="record.keywords" :max-length="250" :auto-size="{minRows:3,maxRows:5}" show-word-limit />
    </a-form-item>

    <a-form-item field="views" :label="$t('views')">
      <a-input-number class="input" v-model="record.views" :min="0" />
    </a-form-item>
  
    <a-form-item field="create_time" :label="$t('createTime')">
      <a-date-picker class="w-full input" style="background-color: var(--color-bg-5);"  v-model="createTime" value-format="timestamp" show-time @change="(val)=>record.create_time =parseInt(val / 1000)" />
    </a-form-item>

    <a-collapse expand-icon-position="right" :default-active-key="['extends']">
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
  import Tag from "./com/Tag.vue"
  import {useStore} from "@/store/index.js";
  import {useOpenLink} from '@/hooks/utils.js'
  import {useAppendSiteURL} from "@/hooks/app/index.js";
  import SelectCategory from "@/components/data/SelectCategory.vue"


  const record = inject('record')
  const createTime = computed(()=>record.value.create_time*1000)
  const store = useStore()

  function addExtends(){
    if(!record.value.extends) record.value.extends = []
    record.value.extends.push({key:'',value:''})
  }

  const slugURL = computed(()=> useAppendSiteURL(store, store.config.router.article_rule.replace('{slug}', record.value.slug)))

  function formatLabel(val){
    console.log(val)
    return "aaa"
  }

</script>


<style scoped>
  .input{
    background-color: var(--color-bg-5);
  }
  .input_extends{
    border-color: var(--color-border-3);
  }
</style>
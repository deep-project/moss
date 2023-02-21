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
  import {useStore} from "@/store/index.js";
  import {useOpenLink} from '@/hooks/utils.js'
  import {useAppendSiteURL} from "@/hooks/app/index.js";

  const record = inject('record')
  const createTime = computed(()=>record.value.create_time*1000)
  const store = useStore()

  const slugURL = computed(()=>{
    return useAppendSiteURL(store, store.config.router.tag_rule.replace('{slug}', record.value.slug))
  })
</script>
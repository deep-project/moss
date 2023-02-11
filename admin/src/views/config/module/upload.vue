<template>
  <a-spin :loading="initLoading">
    <a-form-item :label="$t('domain')">
      <a-input class="w-64" v-model="data.domain" placeholder="/upload/ or https://img.xxx.com" />
    </a-form-item>
    <a-form-item :label="$t('path')">
      <a-radio-group v-model="data.path_format" name="radiogroup">
        <a-space>
          <a-radio key="date" value="date">date</a-radio>
          <a-radio key="hashDate" value="hashDate">hash date</a-radio>
          <a-radio key="hashName" value="hashName">hash name</a-radio>
        </a-space>
      </a-radio-group>
    </a-form-item>
    <a-form-item :label="$t('name')">
      <a-radio-group v-model="data.name_format" name="radiogroup">
        <a-space>
          <a-radio key="original" value="">original</a-radio>
          <a-radio key="md5" value="md5">md5</a-radio>
          <a-radio key="uuid" value="uuid">uuid</a-radio>
          <a-radio key="snowflake" value="snowflake">snowflakeID</a-radio>
        </a-space>
      </a-radio-group>
    </a-form-item>
    <Storage :value="data.storage" :label="$t('storage')" :range='getOptions(["local","ftp","s3","b2","cos","oss"])' />
  </a-spin>
</template>

<script setup>

  import {getOptions} from '@/components/storage'
  import Storage from "@/components/storage/Storage.vue";
  import {useRequest} from "vue-request";
  import { uploadInit } from "@/api";
  import {Message} from '@arco-design/web-vue'
  import {t} from "@/locale";
  import {inject} from 'vue'

  const data = inject('data')
  let messageReactive

  const {run:init,loading:initLoading} = useRequest(uploadInit,{
    manual:true,
    onBefore:()=>{
      messageReactive = Message.loading({content:t('init') + '...', duration:0})
    },
    onSuccess:()=>{
      messageReactive.close()
    }})

  const useSaveSuccess = inject('useSaveSuccess')
  useSaveSuccess.value.push(()=>{ init() })
  
</script>
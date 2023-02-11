<template>
  <a-modal v-model:visible="visible"
           modal-class="shadow-lg"
           :mask-closable="false"
           unmount-on-close
           :closable="false"
           :footer="false"
           :esc-to-close="false"
           title-align="start"
           :mask="false"
           :width="320">
    <template #title>
      <div class="w-full flex items-center justify-between">
        <div>{{ $t('createAdmin') }}</div>
        <Locale type="text" />
      </div>
    </template>
    <a-form :model="form" @submit="onSubmit" layout="vertical">
      <a-form-item field="username" hide-label :rules="[{required:true, message:'Username is required'}]">
        <a-input v-model="form.username" :placeholder="$t('username')" />
      </a-form-item>
      <a-form-item field="password" hide-label :rules="[{required:true, message:'Password is required'}]">
        <a-input-password v-model="form.password" :placeholder="$t('password')" allow-clear />
      </a-form-item>
      <a-form-item field="repeatPassword" hide-label :rules="[{required:true, message:'Enter password again'},{ validator: validateRepeatPassword, trigger: 'change' }]">
        <a-input-password v-model="form.repeatPassword" :placeholder="$t('password')" allow-clear />
      </a-form-item>
      <a-button class="mt-2" type="primary" long html-type="submit" :loading="loading">{{ $t('submit') }}</a-button>
    </a-form>
  </a-modal>
</template>


<script setup>
  import {ref} from 'vue'
  import Locale from '@/components/app/Locale.vue'
  import {adminCreate} from "@/api";
  import {useRequest} from "vue-request";
  import {useRouter} from "vue-router"
  import { Message } from '@arco-design/web-vue';

  const router = useRouter()
  const visible = ref(true)

  const form = ref({
    username:"",
    password:"",
    repeatPassword:"",
  })

  function onSubmit({values, errors}){
    if(errors ===undefined) run(form.value)
  }
  function validateRepeatPassword(v,e){
    if(v!==form.value.password) e("Repeat password do not match.")
  }
  const { loading, run } = useRequest(adminCreate, {
    manual: true,
    loadingDelay: 200,
    loadingKeep: 1000,
    onSuccess:(resp)=>{
      if(resp.success) {
        Message.success("create success")
        visible.value = false
        router.push({name:"login"})
      }
    },
  })

</script>
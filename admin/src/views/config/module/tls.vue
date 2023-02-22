<template>



  <a-form-item :label="$t('enable')">
    <a-switch type="round" v-model="data.enable" />
  </a-form-item>

  <template v-if="data.enable">
    <a-form-item label="KEY" field="key_pem" :rules="[{required:data.enable, message:$t('message.required',['key'])}]">
      <a-textarea v-model="data.key_pem"  style="width:650px" :auto-size="{ minRows: 10, maxRows: 10 }"/>
    </a-form-item>

    <a-form-item label="Cert(pem)" field="cert_pem" :rules="[{required:data.enable, message:$t('message.required',['cert'])}]">
      <a-textarea v-model="data.cert_pem"  style="width:650px" :auto-size="{ minRows: 10, maxRows: 10 }"/>
    </a-form-item>

    <a-form-item label="Listen Addr">
      <a-input v-model="data.listen_addr" style="width:240px" placeholder=":443" />
    </a-form-item>

  </template>

  <a-alert type="warning">{{$t('needRestartApp')}}</a-alert>

</template>

<script setup>
  import {inject,ref} from 'vue'
  import Duration from '@/components/utils/Duration.vue'

  const data = inject('data')

</script>
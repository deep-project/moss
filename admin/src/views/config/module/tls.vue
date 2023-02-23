<template>



  <a-form-item :label="$t('enable')">
    <a-switch type="round" v-model="data.enable" />
  </a-form-item>

  <template v-if="data.enable">
    <a-form-item label="Key" field="key_pem" :rules="[{required:data.enable, message:$t('message.required',['key'])}]">
      <a-textarea v-model="data.key_pem"  style="width:650px" :auto-size="{ minRows: 10, maxRows: 10 }"/>
    </a-form-item>

    <a-form-item label="Cert(pem)" field="cert_pem" :rules="[{required:data.enable, message:$t('message.required',['cert'])}]">
      <a-textarea v-model="data.cert_pem"  style="width:650px" :auto-size="{ minRows: 10, maxRows: 10 }"/>
    </a-form-item>

    <a-form-item label="Port" help="default 443">
      <a-input-number v-model="data.port" style="width:140px" :min="0" />
    </a-form-item>

    <a-form-item label="Force Https">
      <a-space>
        <a-switch type="round" v-model="data.force_https" />
        <a-radio-group v-if="data.force_https" v-model="data.redirect_status">
          <a-radio :value="301">301</a-radio>
          <a-radio :value="302">302</a-radio>
        </a-radio-group>
      </a-space>
    </a-form-item>

  </template>

  <a-divider />

  <a-alert type="warning">{{$t('needRestartApp')}}</a-alert>

</template>

<script setup>
  import {inject,ref} from 'vue'
  import Duration from '@/components/utils/Duration.vue'

  const data = inject('data')

</script>
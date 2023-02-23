<template>
  <a-form-item :label="$t('username')" field="username" hide-asterisk :rules="[{required:true, message:$t('message.required',[$t('username')])}]">
    <a-input v-model="data.username" class="w-64" />
  </a-form-item>
  <a-form-item :label="$t('password')" field="password" :rules="[{ validator: validateRepeatPassword, trigger: 'change' }]">
    <a-space class="w-64" direction="vertical" fill>
      <a-input-password v-model="data.password" placeholder="********" />
      <a-input-password v-model="rePassword" placeholder="Enter again"  />
    </a-space>
  </a-form-item>
  <a-form-item :label="$t('expire')">
    <Duration :data="data.login_expire" />
  </a-form-item>

  <a-divider />

</template>

<script setup>
  import {inject,ref} from 'vue'
  import Duration from '@/components/utils/Duration.vue'

  const data = inject('data')
  data.value.password = "" // 置空密码
  const rePassword = ref()

  function validateRepeatPassword(v,e){
    if(v!==rePassword.value) e("Repeat password do not match.")
  }

</script>
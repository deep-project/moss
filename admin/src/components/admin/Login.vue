<template>
  <a-modal v-model:visible="visible"
           modal-class="shadow-lg"
           :mask-closable="false"
           unmount-on-close
           :closable="false"
           :footer="false"
           :esc-to-close="false"
           title-align="start"
           :mask="showMask"
           :mask-style="{backdropFilter: 'blur(2px)'}"
           :width="320">
    <template #title>
      <div class="w-full flex items-center justify-between">
        <div>{{ $t('login') }}</div>
        <Locale v-if="showLangSelect" type="text" />
      </div>
    </template>
    <a-form :model="form" @submit="onSubmit" layout="vertical">
      <a-form-item field="username" hide-label :rules="[{required:true, message:t('message.required',[t('username')])}]">
        <a-input v-model="form.username" :placeholder="$t('username')" />
      </a-form-item>
      <a-form-item field="password" hide-label :rules="[{required:true, message:t('message.required',[t('password')])}]">
        <a-input-password v-model="form.password" :placeholder="$t('password')" allow-clear />
      </a-form-item>
      <div class="grid grid-cols-2 gap-4">
        <a-form-item field="captcha" hide-label :rules="[{required:true, message:t('message.required',[t('captcha')])}]">
          <a-input v-model="form.captcha" :placeholder="$t('captcha')" />
        </a-form-item>
        <div class="captcha">
          <a-spin class="w-full h-full" :class="{'cursor-wait':captchaLoading}" :loading="captchaLoading">
            <img :src="captchaBase64" @click="onGetCaptcha" />
          </a-spin>
        </div>
      </div>
      <a-button class="mt-2" type="primary" long html-type="submit" :loading="loading">{{ $t('submit') }}</a-button>
    </a-form>
  </a-modal>
</template>


<script setup>
  import {ref} from 'vue'
  import Locale from '@/components/app/Locale.vue'
  import {adminCaptcha, adminExists, adminLogin} from "@/api/index.js";
  import {useRequest} from "vue-request";
  import {useStore} from "@/store/index.js";
  import {useRouter} from "vue-router"
  import {Message} from "@arco-design/web-vue";
  import {t} from "@/locale";

  defineProps({showLangSelect:Boolean, showMask:Boolean})
  const emit = defineEmits(['success'])
  const store = useStore()
  const router = useRouter()
  const visible = ref(true)
  const captchaBase64 = ref("")

  const form = ref({
    username:"",
    password:"",
    captcha:"",
    captchaID:"",
  })

  useRequest(adminExists,{
    onSuccess:(resp)=>{
      if(!resp.data) router.push({name:"createAdmin"})
    },
  })

  const onSubmit = ({values, errors}) => {
    if(errors ===undefined) onLogin(form.value)
  }

  const { loading, run:onLogin } = useRequest(adminLogin, {
    manual: true,
    loadingDelay: 200,
    onSuccess:(resp)=>{
      if(!resp.success){
        if(resp.message !== "captcha is wrong") onGetCaptcha()
        return
      }
      visible.value = false
      store.token = resp.data
      Message.success(t('message.success',[t('login')]))
      emit('success')
    },
  })


  const { loading:captchaLoading, run:onGetCaptcha } = useRequest(adminCaptcha, {
    debounceInterval: 100,
    onSuccess:(resp)=>{
      form.value.captchaID = resp.data.id
      captchaBase64.value = resp.data.base64
    },
  })

</script>

<style scoped>
.captcha{
  position: relative;
  top:-5px;
  height:42px;
}
.captcha img{
  cursor: pointer;
  height: 100%;
  max-width: 100%;
  margin:auto;
  -webkit-tap-highlight-color: transparent;
  user-select:text;
}
</style>

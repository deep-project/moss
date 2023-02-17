<template>

    <div class="p-5">
      {{$t('enable')}}: <a-switch type="round" v-model="data.enable" />
    </div>

    <a-tabs v-if="data.enable" type="card-gutter" destroy-on-hide lazy-load>
      <a-tab-pane key="options" :title="$t('options')">
        <div class="p-3"><Options /></div>
      </a-tab-pane>
      <a-tab-pane key="storage" :title="$t('storage')">
        <div class="p-3"><Storage /></div>
      </a-tab-pane>
<!--      <a-tab-pane key="more" :title="$t('more')">-->
<!--        <div class="p-3"><More /></div>-->
<!--      </a-tab-pane>-->
    </a-tabs>

</template>

<script setup>

  import {useRequest} from "vue-request";
  import {cacheInit} from "@/api";
  import {Message} from "@arco-design/web-vue";
  import {t} from "@/locale";
  import Options from './cache/Options.vue'
  import More from "./cache/More.vue";
  import Storage from "./cache/Storage.vue";
  import {inject} from 'vue'


  const data = inject('data')
  let messageReactive

  const {run:initCache} = useRequest(cacheInit,{
    manual:true,
    onBefore:()=>{
      messageReactive = Message.loading({content:t('init') + '...', duration:0})
    },
    onSuccess:(resp)=>{
      if(!resp.success) Message.error(t('message.failed',[t('init')]))
      messageReactive.close()
    }})

  const useSaveSuccess = inject('useSaveSuccess')
  useSaveSuccess.value.push(()=>{
    initCache()
  })
  
</script>
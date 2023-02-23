<template>
  <a-tabs type="rounded" lazy-load destroy-on-hide>
    <a-tab-pane key="rules" :title="$t('rules')">
      <Rules />
    </a-tab-pane>
    <a-tab-pane key="options" :title="$t('options')">
      <Options />
    </a-tab-pane>

    <a-tab-pane key="pprof" title="pprof">
      <Pprof />
    </a-tab-pane>

  </a-tabs>

  <a-divider />
  
</template>

<script setup>
  import Rules from './router/rules.vue'
  import Options from './router/options.vue'
  import Pprof from './router/pprof.vue'
  
  import {inject} from 'vue'
  import {routerReload} from "@/api";
  import {useRequest} from 'vue-request'
  import {Message} from '@arco-design/web-vue'

  const useSaveSuccess = inject('useSaveSuccess')
  const data = inject('data')
  const oldAdminPath = data.value.admin_path

  useSaveSuccess.value.push(()=>{
    reload()
  })


  const { run:reload } = useRequest(routerReload,{
    manual:true,
    onSuccess:()=>{
      if(oldAdminPath !== data.value.admin_path && process.env.NODE_ENV === "production"){
        Message.loading({content:'loading...', duration:3000})
      setTimeout(()=>{
        self.location = data.value.admin_path.indexOf("/") === 0 ? data.value.admin_path : "/" + data.value.admin_dir
      },3000)
      }
    }
  })
</script>
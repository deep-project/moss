<template>
  <a-card class="w-full" :title="$t(id)" :bordered="false">
    <a-skeleton v-if="loading">
      <a-space direction="vertical" :style="{width:'100%'}" size="large">
        <a-skeleton-line :rows="6" :widths="['80%','90%','92%','76%','82%']" />
      </a-space>
    </a-skeleton>
    <a-form v-else :model="data" auto-label-width @submit="handleSubmit" :layout="store.isMobile ? 'vertical':'horizontal'">
      <component v-bind:is="currentComponent"></component>
      <a-form-item v-if="showBtn" class="mt-2">
        <a-button type="primary" html-type="submit" :loading="loadingSave">{{$t('save')}}</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>


<script setup>

  import {useStore} from "@/store/index.js";
  import {useRoute} from "vue-router";
  import {useRequest} from "vue-request";
  import {configGet,configPost,configList} from "@/api/index.js";
  import {ref, shallowRef, defineAsyncComponent, provide, watch, onMounted} from 'vue'
  import {Message} from "@arco-design/web-vue";
  import {t} from '@/locale'

  const store = useStore()
  const route = useRoute()
  const id = route.params.id.toString()
  const currentComponent = shallowRef()
  const viteComponents = import.meta.glob("./module/*.vue");
  const loading = ref(true)

  const data = ref({})
  provide('data', data)

  const showBtn = ref(true)
  provide('showBtn', showBtn)

  // hooks
  const useSaveBefore = ref([])
  const useSaveSuccess = ref([])    // 保存成功之后的钩子集合
  provide('useSaveBefore', useSaveBefore)
  provide('useSaveSuccess', useSaveSuccess)

  useRequest(configGet, {
    defaultParams:[id],
    onBefore:()=>{},
    onSuccess:(resp)=>{
      currentComponent.value = defineAsyncComponent(viteComponents['./module/'+id+'.vue'])
      data.value = resp.data
      setTimeout(()=>{loading.value = false}, 200)
    }
  });


  const {run:runSave,loading:loadingSave} =useRequest(configPost, {
    manual:true,
    loadingKeep:200,
    onSuccess:(resp)=>{
      if(!resp.success) return
      store.config[id] = data
      for(let item of useSaveSuccess.value) item(data)
      initConfig()
      Message.success(t('message.success',[t('save')]))
    }
  });

  const {run:initConfig} = useRequest(configList, {manual:true, onSuccess:store.initConfig })


  function handleSubmit({values, errors}){
    if(errors!==undefined) return
    for(let item of useSaveBefore.value){
      if (item(data) === false) return
    }
    runSave(id, data.value)
  }

</script>

<style>
.w-64 {
  width:256px !important;
}
.w-32 {
  width:128px !important;
}
</style>
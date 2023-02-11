<template>
  <a-space>
      <a-switch type="round" v-model="data.enable" />
      <template v-if="data.enable">
        <Duration :data="data.ttl" />
        <a-button size="mini" :disabled="cacheData.active_driver==='memcached'" type="text" @click="clear(name)">{{$t('clear')}}</a-button>
        <a-button size="mini" v-if="cacheData.active_driver==='memcached' && name === 'index'" type="text" @click="clear('index')">{{$t('clearAll')}}</a-button>
      </template>
  </a-space>
</template>

<script setup>

  import Duration from "@/components/utils/Duration.vue";
  import {useRequest} from "vue-request";
  import {cacheClear} from "@/api";
  import {inject} from 'vue'
  import {Message} from '@arco-design/web-vue'
  import {t} from '@/locale'

  defineProps({data:Object,name:String})

  const cacheData = inject('data')
  const { run:clear } = useRequest(cacheClear,{
    manual:true,
    onSuccess:(resp)=>{
      if(!resp.success) return
      Message.success(t('message.success',[t('clear')]))
    }
  })

</script>

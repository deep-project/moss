<template>
  <a-table :columns="columns" :data="list" :loading="loadingList" :pagination="false" :bordered="false">
    <template #id="{ record,rowIndex,column }">
      <a-tooltip v-if="record.about">
        <icon-info-circle class="ml-2 opacity-60" />
        <template #content>{{record.about}}</template>
      </a-tooltip>
      <icon-info-circle v-else class="ml-2 opacity-20" />
      {{$t(record.id)}}
    </template>

    <template #action="{ record,rowIndex,column }">
      <a-button :disabled="record.hide_logs"  size="small" type="text" @click="onViewLog(record)"><template #icon><icon-bookmark :size="18" /></template></a-button>
      <a-button :disabled="record.no_options" size="small" type="text" @click="onOptions(record)"><template #icon><icon-settings :size="18" /></template></a-button>
      <a-button :disabled="!record.run_enable" size="small" type="text" @click="run(record.id)" :loading="loadingRunObj[record.id]">
        <template #icon><icon-play-circle :size="18" /></template>
      </a-button>
    </template>

    <template #cron="{ record,rowIndex,column }">
      <Cron :record="record" />
    </template>

    <template #time="{ record,rowIndex,column }">
      <a-tooltip v-if="record.run_enable && record[column.dataIndex] > 0" :content="useDateFormat(record[column.dataIndex]*1000, 'YYYY-MM-DD HH:mm:ss').value">
        <n-time :time="record[column.dataIndex]*1000" :to="Date.now()" type="relative" />
      </a-tooltip>

      <span v-else> - </span>
    </template>

    <template #runDuration="{ record,rowIndex,column }">
      <template v-if="record.run_enable">
        <a-tag v-if="record.run_duration > 1000*60" color="purple">{{(record.run_duration/1000/60).toFixed(0)}} minutes</a-tag>
        <a-tag v-else-if="record.run_duration > 1000" color="purple">{{(record.run_duration/1000).toFixed(2)}} s</a-tag>
        <a-tag v-else-if="record.run_duration > 500" color="purple">{{record.run_duration}} ms</a-tag>
        <a-tag v-else-if="record.run_duration > 100" color="orangered">{{record.run_duration}} ms</a-tag>
        <a-tag v-else color="green">{{record.run_duration}} ms</a-tag>
      </template>
      <span v-else> - </span>
    </template>

    <template #runCount="{ record,rowIndex,column }">
      <span v-if="record.run_enable">{{record.run_count}}</span>
      <span v-else> - </span>
    </template>

  </a-table>

  <a-modal v-model:visible="visibleOptions" :title="modalTitle" :width="600" title-align="start" @ok="runSaveOptions(currentID,currentOptionsData)" :ok-loading="loadingSaveOptions">
    <a-skeleton animation :widths="[80]" v-if="loadingGetOptions">
      <a-space direction="vertical" :style="{width:'100%'}" size="large">
        <a-skeleton-line :rows="5" />
      </a-space>
    </a-skeleton>
    <a-form v-else :model="currentOptionsData" auto-label-width :layout="store.isMobile ? 'vertical':'horizontal'">
      <component v-bind:is="currentOptionsComponent"></component>
    </a-form>
  </a-modal>

  <a-modal v-model:visible="visibleCronExpExample" :title="$t('example')" :width="600" simple :footer="false">
    <CronExpExample />
  </a-modal>

  <a-modal v-model:visible="visibleLog" width="96%" :modal-style="{height:'96%'}" simple :footer="false" modal-class="logModal" unmount-on-close>
    <Log />
  </a-modal>

</template>


<script setup>
  import {useRequest} from "vue-request";
  import CronExpExample from "@/views/plugin/CronExpExample.vue";
  import { NTime } from 'naive-ui'
  import {
    pluginList,
    pluginOptions,
    pluginRun,
    pluginSaveOptions
  } from "@/api/index.js";
  import {computed, defineAsyncComponent, provide, ref, shallowRef} from 'vue'
  import {columns} from './index.js'
  import {useStore} from "@/store/index.js";
  import {Message} from '@arco-design/web-vue'
  import {t} from '@/locale'
  import Cron from "./Cron.vue";
  import Log from "./Log.vue";
  import { useDateFormat } from '@vueuse/core'

  const store = useStore()
  const visibleOptions = ref(false)
  const currentID = ref('')
  const currentOptionsData = ref({})
  const currentOptionsComponent = shallowRef()
  const optionsComponents = import.meta.glob("./options/*.vue");
  const list = ref([])
  const visibleCronExpExample = ref(false)
  const visibleLog = ref(false)
  const modalTitle = computed(()=>{
    if(!currentID.value) return
    return t(currentID.value)
  })


  provide("currentID", currentID)
  provide("options", currentOptionsData)
  provide("visibleCronExpExample", visibleCronExpExample)
  provide("visibleLog", visibleLog)

  const { loading:loadingList } = useRequest(pluginList,{onSuccess:(resp)=>{ list.value = resp}})
  const { run:runGetOptions, loading:loadingGetOptions } = useRequest(pluginOptions,{
    manual:true,
    loadingKeep: 400,
    onBefore:()=>{visibleOptions.value=true},
    onSuccess:(resp,[id])=>{
      currentOptionsData.value = resp
      currentOptionsComponent.value = defineAsyncComponent(optionsComponents['./options/'+id+'.vue'])
  }})
  const { run:runSaveOptions, loading:loadingSaveOptions } = useRequest(pluginSaveOptions,{manual:true,onSuccess:(resp)=>{ resp.success ? Message.success(t('message.success',[t('save')])):'' }})
  const loadingRunObj = ref({})
  const { run, loading:loadingRun } = useRequest(pluginRun,{
    manual:true,
    onBefore:([id])=>{  loadingRunObj.value[id] = true },
    onSuccess:(resp,[id])=>{
      loadingRunObj.value[id] = false
      resp.success ? Message.success(t('message.success',[t('run')])):'' }
  })


  function onOptions(record){
    currentID.value=record.id
    runGetOptions(record.id)
  }


  function onViewLog(record){
    currentID.value = record.id
    visibleLog.value = true
  }


</script>


<style>
.logModal>.arco-modal-header{
  display: none;
}
</style>
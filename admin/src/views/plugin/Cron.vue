<template>

  <div class="h-full flex items-center gap-1">
    <a-switch :disabled="!record.cron_enable" size="mini" v-model="record.cron_start" @change="onChangeCronStart(record)" :loading="loadingCronStart || loadingCronStop" />
    <template v-if="record.cron_enable">
      <div style="width: 120px">
        <a-input v-if="record.editCronExp" size="mini" v-model="record.cron_exp" :class="{readOnly:!record.editCronExp}" style="padding-right: 6px">
          <template #suffix><icon-question-circle @click="visibleCronExpExample = true" class="cursor-pointer" /></template>
        </a-input>
        <span v-else class="text-xs text-gray-500" style="padding:0 13px; line-height: 24px">{{record.cron_exp}}</span>
      </div>
      <div style="color:rgb(var(--arcoblue-6))">
        <icon-edit v-if="!record.editCronExp" @click="record.editCronExp = true" class="cursor-pointer opacity-80 hover:opacity-100" />
        <template v-else>
          <icon-save @click="updateCronExp(record.id, record.cron_exp, record)" class="cursor-pointer opacity-80 hover:opacity-100 mr-1" />
          <icon-minus-circle @click="record.editCronExp = false" class="cursor-pointer opacity-80 hover:opacity-100 mr-1" />
        </template>
      </div>
    </template>
  </div>

</template>


<script setup>

  import {inject} from "vue";
  import {useRequest} from "vue-request";
  import {pluginCronStop, pluginCronExp, pluginCronStart} from "@/api/index.js";
  import {Message} from "@arco-design/web-vue";
  import {t} from "@/locale/index.js";

  defineProps({record:Object})
  const visibleCronExpExample = inject("visibleCronExpExample")

  const { run:cronStart, loading:loadingCronStart } = useRequest(pluginCronStart, {
    manual:true,
    onSuccess:(resp)=>{
      if(resp.success) Message.success(t('message.success',[t('start')]))
    }
  })

  const { run:cronStop, loading:loadingCronStop } = useRequest(pluginCronStop, {
    manual:true,
    onSuccess:(resp)=>{
      if(resp.success) Message.success(t('message.success',[t('stop')]))
    }
  })

  const { run:updateCronExp } = useRequest(pluginCronExp, {
        manual:true,
        onSuccess:(resp,params)=>{
          if(!resp.success) return
          Message.success(t('message.success',[t('update')]))
          params[2].editCronExp = false
        }
  })

  function onChangeCronStart(record){
    if(record.cron_start) cronStart(record.id)
    else cronStop(record.id)
  }

</script>

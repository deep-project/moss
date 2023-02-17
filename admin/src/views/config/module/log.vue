<template>
  <a-tabs type="card-gutter" animation style="max-width: 600px">
    <a-tab-pane key="app" :title="$t('app')">
      <Item :data="data.app" />
    </a-tab-pane>
    <a-tab-pane key="sql" :title="$t('sql')">
      <Item :data="data.sql" />
    </a-tab-pane>
    <a-tab-pane key="slow_sql" :title="$t('slow_sql')">
      <Item :data="data.slow_sql" />
    </a-tab-pane>
    <a-tab-pane key="visitor" :title="$t('visitor')">
      <Item :data="data.visitor" />
    </a-tab-pane>
    <a-tab-pane key="spider" :title="$t('spider')">
      <Item :data="data.spider" />
    </a-tab-pane>
    <a-tab-pane key="plugin" :title="$t('plugin')">
      <Item :data="data.plugin" />
    </a-tab-pane>
    <a-tab-pane key="more" :title="$t('more')">
      <More />
    </a-tab-pane>
  </a-tabs>
</template>

<script setup>

  import Item from "./log/Item.vue";
  import {useRequest} from "vue-request";
  import {logInit} from "@/api";
  import {t} from "@/locale";
  import {Message} from '@arco-design/web-vue'
  import {inject} from 'vue'
  import More from './log/more.vue'

  const data = inject('data')
  let messageReactive

  const {run:init} = useRequest(logInit,{
    manual:true,
    onBefore:()=>{
      messageReactive = Message.loading({content:t('init') + '...', duration: 0  })
    },
    onSuccess:()=>{
      messageReactive.close()
    }})

  const useSaveSuccess = inject('useSaveSuccess')
  useSaveSuccess.value.push(()=>{ init() })

</script>
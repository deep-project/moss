<template>
  <a-card class="w-full h-full" size="small" :bordered="false" :body-style="{padding:'0'}" :header-style="{border:'none',padding: '8px 0'}">
    <template #title>
      {{$t('logs')}} <span class="text-sm text-gray-500">{{$t(id)}}</span>
    </template>
    <template #extra>
      <a-space>
        <a-input-number size="small" v-model="limit" :style="{width:'80px'}" :min="1" @keyup.enter="refresh" />
        <a-button size="small" @click="refresh" type="outline">{{$t('refresh')}}</a-button>
        <a-button size="small" @click="loadMore" type="outline">{{$t('more')}}</a-button>
        <a-button size="small" @click="visibleLog = false" type="outline">{{$t('close')}}</a-button>

      </a-space>
    </template>
    <a-table size="mini"
             :columns="columns"
             :data="dataList"
             :pagination="false"
             :scroll="{y:tableHeight}"
             :virtual-list-props="dataList.length > 1 ? {height:tableHeight,buffer:30} : undefined"
             :loading="loading || loadingMore"
             style="--border-radius-medium:0">

      <template #th>
        <th style="color:var(--color-text-1);background-color: var(--color-fill-2)"></th>
      </template>
      <template #td>
        <td style="color:var(--color-text-1)"></td>
      </template>

      <template #detail="{ record,rowIndex,column }">
        <div class="flex items-center"><a-link @click="visibleDetail=true;detailData = record"><icon-eye style="color: #6b7f94" /></a-link></div>
      </template>

      <template #time="{ record,rowIndex,column }">
        <n-time :time="record.time" :to="Date.now()" type="relative" />
      </template>

      <template #more="{ record,rowIndex,column }">
        <div v-if="Object.keys(record[column.dataIndex]).length > 0">
          {{JSON.stringify(record[column.dataIndex]).substring(0,130)}}
        </div>
      </template>

      <template #level="{ record,rowIndex,column }">
        <a-tag size="mini" v-if="record.level ==='info'" color="blue">{{record.level}}</a-tag>
        <a-tag size="mini" v-else-if="record.level ==='debug'" color="cyan">{{record.level}}</a-tag>
        <a-tag size="mini" v-else color="red">{{record.level}}</a-tag>
      </template>

    </a-table>
  </a-card>

  <a-modal v-model:visible="visibleDetail" title="Log detail" title-align="start"
           :modal-style="{width:'auto',height:'92%',minWidth:'600px',maxWidth:'80%'}"
           :body-style="{height: 'calc(100% - 52px)'}"
           :footer="false">
    <a-descriptions :column="1" bordered>
      <a-descriptions-item label="time"><n-time :time="detailData.time" /></a-descriptions-item>
      <a-descriptions-item label="level">{{detailData.level}}</a-descriptions-item>
      <a-descriptions-item label="file">{{detailData.file}}</a-descriptions-item>
      <a-descriptions-item label="msg">{{detailData.msg}}</a-descriptions-item>
      <template v-for="(val,key) in detailData.more">
        <a-descriptions-item :label="key"><code class="font-mono text-gray-500 whitespace-pre-line">{{val}}</code></a-descriptions-item>
      </template>
    </a-descriptions>
  </a-modal>

</template>

<script setup>
  import {computed, inject, ref} from "vue";
  import {useLoadMore} from "vue-request";
  import {pluginLogList} from "@/api/index.js";
  import {useDeepCopy} from "@/hooks/utils.js";
  import { NTime } from 'naive-ui'

  const id = inject("currentID")
  const visibleLog = inject("visibleLog")

  const limit = ref(100)
  const tableHeight = computed(()=> 'calc(90vh - 30px)')
  const visibleDetail= ref(false)
  const detailData = ref({})

  const columns = [
    {
      width: 30,
      slotName:'detail'
    },
    {
      title: 'time',
      dataIndex: 'time',
      slotName:'time',
      width: 120,
    },
    {
      title: 'file',
      dataIndex: 'file',
      ellipsis:true,
      width: 300,
    },
    {
      title: 'level',
      dataIndex: 'level',
      slotName:'level',
      width: 100,
    },
    {
      title: 'message',
      dataIndex: 'msg',
      ellipsis:true,
      width: 300,
    },
    {
      title: 'more',
      dataIndex: 'more',
      slotName:'more',
      width: 400,
    },
  ]


  async function queryData(d){
    const _page = d?.page ? d.page + 1 : 1;
    const results = await pluginLogList(id.value, {page:_page, limit: limit.value})
    return {
      list: useFormatLogList(results),
      page: _page,
    }
  }

  const { loading, loadingMore, dataList, loadMore, refresh } = useLoadMore(queryData,{});

  function useFormatLogList(list){
    list = list.reverse()
    for(let k in list) {
      let _a = useDeepCopy(list[k])
      delete _a.level
      delete _a.time
      delete _a.file
      delete _a.msg
      list[k].more = _a
    }
    return list
  }

</script>
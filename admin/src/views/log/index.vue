<template>
  <a-card class="w-full" size="small" :title="$t(id)" :bordered="false" :body-style="{padding:'0'}" :header-style="{border:'none',padding: '8px 0'}">
    <template #extra>
      <a-space>
        <a-input-number size="small" v-model="limit" :style="{width:'80px'}" :min="1" @keyup.enter="onRefresh" />
        <a-button size="small" @click="onRefresh" type="outline">{{$t('refresh')}}</a-button>
        <a-button size="small" @click="loadMore" type="outline">{{$t('more')}}</a-button>
      </a-space>
    </template>
    <a-table size="mini" ref="tableRef" :bordered="false"
             :columns="columns"
             :data="dataList"
             :pagination="false"
             :scroll="{y:tableHeight, x:'100%'}"
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

      <template #take="{ record,rowIndex,column }">
        <a-tag size="mini" v-if="record.take < 100" color="blue">{{record.take.toFixed(2)}} ms</a-tag>
        <a-tag size="mini" v-else-if="record.take < 200" color="orange">{{parseInt(record.take)}} ms</a-tag>
        <a-tag size="mini" v-else-if="record.take < 500" color="purple">{{parseInt(record.take)}} ms</a-tag>
        <a-tag size="mini" v-else color="red">{{parseInt(record.take)}} ms</a-tag>
      </template>

      <template #status="{ record,rowIndex,column }">
        <a-tag size="mini" v-if="record.status >=500" color="red">{{record.status}}</a-tag>
        <a-tag size="mini" v-else-if="record.status >=400" color="orange">{{record.status}}</a-tag>
        <a-tag size="mini" v-else-if="record.status >=300" color="blue">{{record.status}}</a-tag>
        <a-tag size="mini" v-else color="green">{{record.status}}</a-tag>
      </template>

      <template #url="{ record,rowIndex,column }">
        <a class="hover:underline cursor-pointer hover:text-blue-400" @click="useOpenLink(record[column.dataIndex])">{{record[column.dataIndex]}}</a>
      </template>

      <template #level="{ record,rowIndex,column }">
        <a-tag size="mini" v-if="record.level ==='info'" color="blue">{{record.level}}</a-tag>
        <a-tag size="mini" v-else-if="record.level ==='debug'" color="cyan">{{record.level}}</a-tag>
        <a-tag size="mini" v-else color="red">{{record.level}}</a-tag>
      </template>

      <template #rows="{ record,rowIndex,column }">
        <span v-if="record.rows > -1">{{record.rows}}</span>
        <span v-else> - </span>
      </template>



      <template #more="{ record,rowIndex,column }">
        <div v-if="Object.keys(record[column.dataIndex]).length > 0">
          {{JSON.stringify(record[column.dataIndex]).substring(0,130)}}
        </div>
      </template>

    </a-table>
  </a-card>

  <a-modal v-model:visible="visibleDetail" title="Log detail" title-align="start"
           :modal-style="{width:'auto',height:'92%',minWidth:'600px',maxWidth:'80%'}"
           :body-style="{height: 'calc(100% - 52px)'}"
           :footer="false">
    <DescAPP v-if="id === 'app'" :data="detailData" />
    <DescSQL v-if="id === 'sql' || id === 'slow_sql'" :data="detailData" />
    <DescHTTP v-if="id === 'visitor' || id === 'spider'" :data="detailData" />
  </a-modal>

</template>


<script setup>
  import {useRoute} from "vue-router";
  import {useStore} from "@/store";
  import {useStorage} from "@vueuse/core";
  import {ref,computed} from "vue";
  import {logRead} from '@/api';
  import {useLoadMore} from 'vue-request';
  import {useDeepCopy,useOpenLink} from '@/hooks/utils'
  import {columnAPP, columnSQL, columnHTTP, columnSpider} from './columns'
  import { NTime } from 'naive-ui'
  import DescAPP from './descriptions/app.vue'
  import DescHTTP from './descriptions/http.vue'
  import DescSQL from './descriptions/sql.vue'

  const store = useStore()
  const route = useRoute()
  const id = route.params.id.toString()
  const limit = useStorage("log_limit", 100)
  const tableRef = ref()
  const tableHeight = computed(()=> store.mainHeight - 66)
  const visibleDetail= ref(false)
  const detailData = ref({})

  const columns = computed(()=>{
    switch (id){
      case "app":return columnAPP
      case "sql":return columnSQL;
      case "slow_sql":return columnSQL;
      case "visitor":return columnHTTP;
      case "spider":return columnSpider();
    }
  })

  async function queryData(d){
    const _page = d?.page ? d.page + 1 : 1;
    const results = await logRead(id, {page:_page, limit: limit.value})
    return {
      list: useFormatLogList(results),
      page: _page,
    }
  }

  const { loading, loadingMore, dataList, loadMore, refresh } = useLoadMore(queryData,{});

  function onRefresh(){
    refresh()
    //tableRef.value.scroll({y:0, top: 0 })
  }

  function getTime(t){
    let n = new Date().getTime()
    // console.log("t-",t)
    // console.log("n-",n)
    // n:1670479745578
    // t:1670479745389.976
    if(t > n) return n
    return parseInt(t)
  }


  // 格式化日志列表
  function useFormatLogList(list){
    list = list.reverse()
    if(id!=='app') return list
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

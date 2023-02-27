<template>
  <a-card class="w-full h-full" :title="$t('system')" :bordered="false">
      <a-grid :cols="{ xs: 3, sm: 4, md: 5, lg:6, xl:7, xxl:8 }" :colGap="12" :rowGap="18" class="text-center">
  <!--      <a-grid-item v-if="loadDec > -1">-->
  <!--        <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="dec(load)" status='warning' :color="color" />-->
  <!--        <div class="title">{{$t('load')}}</div>-->
  <!--      </a-grid-item>-->
        <a-grid-item>
            <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="dec(cpu)" status='warning' :color="color" />
            <div class="title">cpu</div>
        </a-grid-item>
        <a-grid-item>
            <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="dec(memory)" status='warning' :color="color" />
            <div class="title">{{$t('memory')}}</div>
        </a-grid-item>

        <a-grid-item v-for="(item,index) in disks">
            <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="dec(item)" status='warning' :color="color" />
            <div class="title">{{$t('disk') + (disks.length > 1 ? ' '+(index+1):'')}}</div>
        </a-grid-item>

      </a-grid>
    </a-card>
</template>



<script setup>

  import {useRequest} from "vue-request";
  import {dashboardData} from "@/api/index.js";
  import {useStore} from "@/store/index.js";
  import {computed, ref} from "vue";

  const store = useStore()
  const color = 'rgb(var(--primary-6))'

  //const { data:load } = useRequest(dashboardData, {defaultParams:['systemLoad'], pollingInterval: 1000, errorRetryCount: 1})
  const { data:cpu } = useRequest(dashboardData, {defaultParams:['systemCPU'], pollingInterval: 2000, errorRetryCount: 1})
  const { data:memory } = useRequest(dashboardData, {defaultParams:['systemMemory'], pollingInterval: 2000, errorRetryCount: 1})
  const { data:disks } = useRequest(dashboardData, {defaultParams:['systemDisk']})

  function dec(val){
    if(val === -1) return -1
    if(!val) return 0
    return Number((val/100).toFixed(2))
  }
</script>

<style scoped>
.title{
  margin-top:2px;
  font-size:12px;
  color:#666666;
}
</style>
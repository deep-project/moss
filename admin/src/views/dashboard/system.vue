<template>
  <a-card class="w-full h-full" :title="$t('system')" :bordered="false">

    <a-grid :cols="{ xs: 4, sm: 4, md: 4 }" class="text-center">
      <a-grid-item>
        <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="loadDec" status='warning' :color="color" />
        <div class="title">load</div>
      </a-grid-item>
      <a-grid-item>
          <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="cpuDec" status='warning' :color="color" />
          <div class="title">cpu</div>
      </a-grid-item>
      <a-grid-item>
          <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="memoryDec" status='warning' :color="color" />
          <div class="title">memory</div>
      </a-grid-item>
      <a-grid-item>
          <a-progress type="circle" :size="store.isMobile ? 'medium':'large'" :percent="diskDec" status='warning' :color="color" />
          <div class="title">disk</div>
      </a-grid-item>
    </a-grid>

  </a-card>
</template>



<script setup>

  import {useRequest} from "vue-request";
  import {dashboard} from "@/api/index.js";
  import {useStore} from "@/store/index.js";
  import {computed} from "vue";

  const store = useStore()
  const color = 'rgb(var(--primary-6))'


  const { data:load } = useRequest(dashboard, {defaultParams:['systemLoad'], pollingInterval: 5000})
  const loadDec = computed(()=>dec(load.value))

  const { data:cpu } = useRequest(dashboard, {defaultParams:['systemCPU'], pollingInterval: 1000})
  const cpuDec = computed(()=>dec(cpu.value))

  const { data:memory } = useRequest(dashboard, {defaultParams:['systemMemory'], pollingInterval: 1000})
  const memoryDec = computed(()=>dec(memory.value))

  const { data:disk } = useRequest(dashboard, {defaultParams:['systemDisk']})
  const diskDec =  computed(()=>dec(disk.value))

  function dec(val){
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
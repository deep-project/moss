<template>
  <SubMenuLayout :width="120">
    <template #sider>
      <a-menu class="menu" v-model:selected-keys="current" @menu-item-click="menuItemClick">
        <router-link v-for="item in ['app','sql','slow_sql','visitor','spider']" :to="{name:'logs-item',params:{id:item}}">
          <a-menu-item :key="item">{{$t(item)}}</a-menu-item>
        </router-link>
      </a-menu>
    </template>
    <router-view v-if="route.params.id" :key="route.params.id.toString()" />
  </SubMenuLayout>
</template>

<script setup>
  import SubMenuLayout from '@/layout/subMenu.vue'
  import {watch,ref} from "vue";
  import {useRoute,useRouter} from "vue-router";
  import {useStorage} from "@vueuse/core";

  const route = useRoute()
  const router = useRouter()
  const current = ref()
  const defLogItem = useStorage("log_page_current",'app')

  function initCurrent(){
   current.value = [route.params.id]
  }

  initCurrent()
  watch(()=>route.name,(val)=>{
    initCurrent()
  })

  if(!route.params.id){
    router.push({ name: 'logs-item', params:{id: defLogItem.value } })
    current.value = [defLogItem.value]
  }

  function menuItemClick(val){
    defLogItem.value = val
  }

</script>

<style scoped>
  .menu{
    background-color: transparent;
  }
  .menu :deep(.arco-menu-inner){
    padding-left: 0;
    padding-right: 6px;
  }
  .menu :deep(.arco-menu-item){
    background-color: transparent;
  }

  .menu :deep(.arco-menu-selected),
  .menu :deep(.arco-menu-item.arco-menu-selected:hover){
    background-color: rgb(var(--arcoblue-1)) !important;
  }
  .menu :deep(.arco-menu-item:hover){
    background-color: #edeff5 !important;
  }
  body[arco-theme="dark"] .menu :deep(.arco-menu-selected),
  body[arco-theme="dark"] .menu :deep(.arco-menu-item.arco-menu-selected:hover){
    background-color: #27324a !important;
  }
  body[arco-theme="dark"] .menu :deep(.arco-menu-item:hover){
    background-color: #2d3138 !important;
  }
</style>
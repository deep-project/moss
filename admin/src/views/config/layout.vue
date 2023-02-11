<template>
  <SubMenuLayout :width="160">
    <template #sider>
      <a-menu class="menu" v-model:selected-keys="current">
        <router-link v-for="(val,key) in store.config" :to="{name:'config-item',params:{id:key}}">
          <a-menu-item :key="key">{{$t(key)}}</a-menu-item>
        </router-link>
      </a-menu>
    </template>
    <router-view v-if="route.params.id" :key="route.params.id.toString()" />
  </SubMenuLayout>
</template>

<script setup>
  import SubMenuLayout from '@/layout/subMenu.vue'
  import {ref, watch} from "vue";
  import {useRoute} from "vue-router";
  import {useStore} from "@/store/index.js";

  const route = useRoute()
  const current = ref()
  const store = useStore()

  function initCurrent(){
    current.value = [route.params.id]
  }
  initCurrent()
  watch(()=>route.name,(val)=>{
    initCurrent()
  })

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
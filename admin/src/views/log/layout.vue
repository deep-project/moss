<template>
  <SubMenuLayout :width="120">
    <template #sider>
      <a-menu class="menu" v-model:selected-keys="current">
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
  import {ref, watch} from "vue";
  import {useRoute} from "vue-router";

  const route = useRoute()
  const current = ref()

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
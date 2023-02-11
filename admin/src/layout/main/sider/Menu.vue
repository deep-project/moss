<template>
  <a-menu
      class="sider_menu"
      :collapsed-width="store.siderCollapsedWidth"
      v-model:selected-keys="current"
      @menu-item-click="onClickMenu"
  >
    <router-link v-for="menu in menuList" :to="{name:menu.name}">
      <a-menu-item :key="menu.name">
        <template #icon><component v-bind:is="menu.meta.icon" :size="20" /></template>
        {{ $t(menu.name ) }}
      </a-menu-item>
    </router-link>
  </a-menu>
</template>

<script setup>
  import {useStore} from "@/store";
  import {computed, ref, watch} from 'vue'
  import {useRoute} from "vue-router";
  import routes from '@/router/routes'

  const route = useRoute()
  const store = useStore()
  const current = ref()
  const menuList = computed(()=>{
    for(let i in routes){
      if(routes[i].name === 'index') return routes[i].children
    }
  })

  function initCurrent(){
    current.value = [route.name?.toString().split("-")[0]]
  }

  initCurrent()
  watch(()=>route.name,(val)=>{
    initCurrent()
  })


  function onClickMenu(){
    if(store.isMobile) store.siderCollapsed = true
  }
</script>

<style>
.sider_menu .arco-menu-inner{
  padding:4px 7px !important;
}
.sider_menu.arco-menu-collapsed .arco-menu-icon{
  width: 100%;
  text-align: center;

}
.sider_menu .arco-menu-icon{
  margin-right: 10px !important;
}

.sider_menu.arco-menu-collapsed .arco-menu-item,
.sider_menu.arco-menu-collapsed .arco-menu-pop-header{
  padding: 0 !important;
}
.sider_menu.arco-menu-vertical .arco-menu-item,
.sider_menu.arco-menu-vertical .arco-menu-group-title,
.sider_menu.arco-menu-vertical .arco-menu-pop-header,
.sider_menu.arco-menu-vertical .arco-menu-inline-header{
  line-height: 44px !important;
  margin-bottom: 5px !important;
  border-radius: 4px;
}

.sider_menu.arco-menu-light .arco-menu-item.arco-menu-selected{
  background-color: #e0e9f7 !important;
}
body[arco-theme="dark"] .sider_menu .arco-menu-item.arco-menu-selected{
  background-color: #27324a !important;
}

body:not([arco-theme="dark"]) .sider_menu .arco-menu-pop-header:not(.arco-menu-selected):hover,
body:not([arco-theme="dark"]) .sider_menu .arco-menu-inline-header:not(.arco-menu-selected):hover,
body:not([arco-theme="dark"]) .sider_menu .arco-menu-item:not(.arco-menu-selected):hover{
  background-color: #e8e9f0 !important;
}
</style>

<template>
  <a-layout-sider
      hide-trigger
      collapsible
      breakpoint="md"
      :width="store.siderWidth"
      :collapsed-width="store.siderCollapsedWidth"
      :collapsed="store.siderCollapsed"
      @breakpoint="breakpoint"
      class="h-full z-20 relative"
      :style="{position: store.isMobile ? 'fixed':''}"
  >
    <Logo />
    <Menu />
    <WidthChange />
  </a-layout-sider>
  <div v-if="store.isMobile && !store.siderCollapsed" @click="store.siderCollapsed = true"
       class="fixed z-10 w-screen h-screen bg-gray-800 top-0 left-0 bg-opacity-30" style="backdrop-filter: blur(2px)"></div>
</template>


<script setup>
  import {useStore} from "@/store";
  import Logo from './Logo.vue'
  import Menu from "./Menu.vue";
  import WidthChange from "./WidthChange.vue";

  const store = useStore()

  function breakpoint(e){
    if(e) store.siderCollapsed = true
  }
</script>

<style>
.arco-layout-sider-light{
  box-shadow:none !important;
}
</style>
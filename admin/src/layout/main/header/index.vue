<template>
  <a-layout-header :style="{height:store.headerHeight+'px',lineHeight:store.headerHeight+'px'}">
    <div class="grid grid-cols-12 px-3 h-full overflow-hidden">
      <div class="col-span-2 sm:col-span-6 md:col-span-6 flex items-center h-full overflow-hidden">
        <SiderCollapsed class="mr-2" />
        <Breadcrumb v-if="!store.isMobile" />
      </div>
      <div class="col-span-10 sm:col-span-6 md:col-span-6 flex justify-end items-center">
        <Dark class="mr-2" />
          <Home />
          <Locale />
          <Logout />
      </div>
    </div>
  </a-layout-header>
</template>


<script setup>
  import Locale from "@/components/app/Locale.vue";
  import Dark from "@/components/app/Dark.vue";
  import Logout from "@/components/app/Logout.vue";
  import SiderCollapsed from "./SiderCollapsed.vue";
  import Home from './Home.vue'
  import Breadcrumb from "./Breadcrumb.vue";
  import {useStore} from "@/store/index.js";

  const store =useStore()
</script>
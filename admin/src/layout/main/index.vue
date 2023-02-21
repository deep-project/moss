<template>
  <div class="w-full h-full" :class="{'p-0 xl:p-5':!store.isMobile}">
    <a-layout ref="layout" class="w-full h-full overflow-hidden layout" :class="{'rounded-xl':!store.isMobile}" :style="{backgroundColor:store.dark ? 'var(--color-menu-dark-bg)':''}">
      <Sider />
      <a-layout class="h-full">
        <Header />
        <div class="overflow-auto mx-2" :style="{height:store.mainHeight+'px'}">
            <div class="h-full mx-1"><router-view /></div>
        </div>
      </a-layout>
    </a-layout>
  </div>
  <Login class="login" v-if="showLogin" :show-mask="true"  />
</template>

<script setup>
  import Sider from '@/layout/main/sider/index.vue'
  import Header from '@/layout/main/header/index.vue'
  import Login from '@/components/admin/Login.vue'
  import {useStore} from "@/store/index.js";
  import { useElementSize } from '@vueuse/core'
  import {ref, computed, onMounted, watch} from "vue";
  import {configList} from "@/api/index.js";
  import {useRequest} from "vue-request";
  import {useRoute, useRouter} from 'vue-router'

  const store = useStore()
  const router = useRouter()
  const route = useRoute()

  const showLogin = computed(()=>{
    if(!store.token && route.name === "dashboard"){
      router.push({name:"login"}) // 如果在控制台失去权限，直接跳转到登录页，防止pwa授权后无法更新页面
      return
    }
    return !store.token
  })
  const layout = ref()
  const { height } = useElementSize(layout)

  watch(height,(val)=>{
    store.mainHeight = val - store.headerHeight - 8
  })

  // init config
  useRequest(configList, {onSuccess:store.initConfig})
</script>


<style scoped>
.layout{
  color: var(--color-text-1);
  background-color: var(--color-fill-1);
}

</style>
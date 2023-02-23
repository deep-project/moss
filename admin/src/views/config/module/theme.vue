<template>

  <a-grid :rowGap="20" :colGap="14" :cols="{ xs: 1, sm: 2, md: 3, lg: 4 }" class="w-full mb-4">
    <a-grid-item v-for="theme in themeListData">
      <a-card class="h-full cursor-pointer" size="mini" style="border-width:2px;"
              :class="{current:theme.id === data.current}" :style="{borderColor:theme.id === data.current ? store.color:''}"
              @click="data.current = theme.id">
        <template #header>
          <div class="flex place-items-center">
            <div>{{theme.name ? theme.name : theme.id}}</div>
            <a-badge :value="theme.version" class="ml-2" :color="store.dark ? '#5d6a77':'#b7bfce'" />
          </div>
        </template>
        <template #cover>
            <div class="overflow-hidden image" :style="{'backgroundColor':store.dark ? '#282c34':''}">
              <div v-if="theme.has_screenshot" class="w-full h-full">
                <a-spin :loading="!theme.screenshot" class="w-full h-full" :size="32">
                  <img v-if="theme.screenshot" class="w-full h-full p-1" :src="theme.screenshot" :alt="theme.name"  />
                </a-spin>
              </div>
              <div v-else class="w-full h-full grid justify-items-center items-center">
                <span class="opacity-20 text-gray-500"><icon-image-close :size="120" /></span>
              </div>
            </div>
            <div v-if="theme.id === data.current" class="absolute right-0 top-0 z-5 check" :style="{borderRightColor:store.color}">
              <span class="absolute text-white" style="top:3px;right:-40px"><icon-check-circle :size="20" /></span>
            </div>
        </template>
          <a-descriptions :column="1" size="mini" :value-style="{paddingRight:'0'}">
            <template #title>
                <b>{{theme.name}}</b>
                <a-tag v-if="theme.version" size="mini" class="ml-2" color="arcoblue" style="border-radius: 30px">{{theme.version}}</a-tag>
            </template>
            <a-descriptions-item v-if="theme.author" label="author">{{ theme.author }}</a-descriptions-item>
            <a-descriptions-item v-if="theme.homepage" label="homepage">
              <span @click.stop="useOpenLink" class="hover:underline">{{ theme.homepage }}</span>
            </a-descriptions-item>
            <a-descriptions-item v-if="theme.license" label="license">{{ theme.license }}</a-descriptions-item>
            <a-descriptions-item v-if="theme.description" label="description">{{ theme.description }}</a-descriptions-item>
          </a-descriptions>
      </a-card>
    </a-grid-item>
  </a-grid>

  <a-divider />
  
</template>

<script setup>

  import {useRequest} from "vue-request";
  import {routerReload, themeInit, themeList, themeScreenshot} from "@/api";
  import {useStore} from "@/store";
  import {useOpenLink} from '@/hooks/utils'
  import {inject,ref} from 'vue'

  const data = inject('data')
  const store = useStore()
  const themeListData = ref([])

  const {run:runRouterReload} = useRequest(routerReload,{manual:true})
  const {run:initTheme} = useRequest(themeInit,{ manual:true, onSuccess:()=>{
      runRouterReload()
    }
  })

  const {run:getThemeList} = useRequest(themeList,{
    onSuccess:(data)=>{
      themeListData.value = data
      for(let i=0;i<data.length;i++) {
        useRequest(themeScreenshot,{ defaultParams:[data[i].id], onSuccess:(resp,params)=>{ themeListData.value[i].screenshot = resp } })
      }
    }
  })


  const useSaveSuccess = inject('useSaveSuccess')
  useSaveSuccess.value.push(()=>{ initTheme() })

</script>

<style scoped>
.check{
  width: 0;
  height: 0;
  border-width: 0 40px 40px 0;
  border-style: solid;
  border-color: transparent transparent transparent;
}

.image{
  height: 230px
}
@media (max-width:576px) {
  .image{
    height: 220px
  }
}
@media (min-width:860px) {
  .image{
    height: 240px
  }
}
@media (min-width:1024px) {
  .image{
    height: 260px
  }
}
@media (min-width:1920px) {
  .image{
    height: 320px
  }
}

</style>
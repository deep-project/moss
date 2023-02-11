<template>
  <a-layout-sider
      class="relative"
      collapsible
      hide-trigger
      v-model:collapsed="collapsed"
      :collapsed-width="store.isMobile ? 10:14"
      :width="width"
  >
      <span class="cursor-pointer absolute z-10 top-1/2 rounded-full flex items-center p-1"
            style="left:-8px;margin-top:-11px;color:rgb(var(--arcoblue-5));background-color: var(--color-menu-light-bg)"
            @click="collapsed=!collapsed">
          <icon-left class="opacity-60 hover:opacity-100 transition" style="margin-left: -2px;" v-if="collapsed" :size="16" :stroke-width="5" />
          <icon-right class="opacity-60 hover:opacity-100 transition" style="margin-left: -2px;" v-else :size="16" :stroke-width="5" />
      </span>

      <div class="absolute bottom-1 left-1 z-10 cursor-pointer hover:text-gray-400 transition" v-if="!collapsed && !store.isMobile"
           style="cursor: col-resize;color:var(--color-border-4)" @mousedown="onmousedown($event.pageX)"
           @touchstart="onmousedown($event.touches[0].clientX)"
           @dblclick.native="doubleClick">
        <icon-pause class="select-none opacity-60 hover:opacity-100 transition" />
      </div>
      <div class="overflow-hidden" :class="{hidden:collapsed}">
        <div class="py-3 px-5" :style="{width:width+'px'}"><PostRightContent /></div>
      </div>
  </a-layout-sider>
</template>

<script setup>
  import {inject,ref} from "vue";
  import {useStore} from "@/store/index.js";
  import {useStorage} from "@vueuse/core";
  import PostRightContent from "@/views/article/PostRightContent.vue";

  const store=useStore()
  const collapsed = ref(store.isMobile)
  const defWidth = store.isTablet ? 270:340
  const width = useStorage("article_post_right_width", defWidth)
  if(store.isMobile) width.value = '96%'

  function onmousedown(pageX){
    let isDown = true
    let maxW = store.windowSize.width*0.8
    let minW = 100
    let startW = width.value

    let event = function (newPageX) {
      let w = startW - (newPageX - pageX)
      if(isDown && w > minW && w < maxW) width.value = w
    }

    document.onmousemove = (ev)=> event(ev.pageX)
    document.onmouseup = function (e) {
      isDown = false
      document.onmousemove = null;
      document.onmouseup = null;
    }
    const touchEvent = (e)=> {
      //e.preventDefault()
      event(e.touches[0].pageX)
    }
    document.addEventListener('touchmove', touchEvent)
    document.addEventListener('touchend', function(){
      isDown = false
      document.removeEventListener('touchmove', touchEvent)
    })
  }

  function doubleClick(){
    width.value = defWidth
  }

</script>

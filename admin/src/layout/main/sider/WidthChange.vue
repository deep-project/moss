<template>
  <div class="absolute bottom-1 right-1 z-10 cursor-pointer text-gray-300 hover:text-gray-400 transition"
       :class="{'text-gray-600':store.dark}" v-if="!store.siderCollapsed"
       style="cursor: col-resize" @mousedown="onmousedown($event.pageX)"
       @touchstart="onmousedown($event.touches[0].clientX)"
       @dblclick.native="doubleClick">
    <icon-pause class="select-none" />
  </div>
</template>

<script setup>

 import {useStore} from "@/store/index.js";

 const store = useStore()


 function onmousedown(pageX){
   let isDown = true
   let startW = store.siderWidth
   let maxW = store.windowSize.width*0.8
   let minW = 100

   let event = function (newPageX) {
     let width = startW + (newPageX - pageX)
     if(isDown && width > minW && width < maxW) store.siderWidth = width
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
   store.siderWidth = 180
 }
</script>
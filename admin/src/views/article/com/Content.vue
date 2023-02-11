<template>
  <ContentEditor />
  <div class="absolute z-5 bottom-2 right-3 cursor-pointer opacity-10 hover:opacity-20 hover:text-blue-800 transition"
       :class="{'hover:text-white':store.dark}" @click="visible = true">
    <icon-code-square :size="50" />
  </div>
  <a-modal width="96%" v-model:visible="visible" @cancel="modalClose" unmount-on-close
           modal-class="codeModal"
           :mask-style="{backdropFilter: 'blur(2px)'}"
           :modal-style="{height:'96%',padding:'10px',backgroundColor:store.dark ? '#282c34':'#f5f5f5'}"
           :body-style="{height:'100%',overflow:'hidden'}"
           simple
           :footer="false"
  >
    <ContentHtmlCode ref="codeRef" />
    <div class="cursor-pointer absolute right-1 top-1 opacity-10 hover:opacity-20 hover:text-blue-800 transition"
         :class="{'opacity-20 hover:text-white':store.dark}" @click="modalClose">
      <icon-close-circle :size="40" />
    </div>
  </a-modal>
</template>
<script setup>
  import ContentEditor from "@/views/article/com/ContentEditor.vue";
  import ContentHtmlCode from './ContentHtmlCode.vue'
  import {ref} from 'vue'
  import {useStore} from "@/store/index.js";

  const store = useStore()
  const visible = ref(false)
  const codeRef = ref(null)

  function modalClose(){
    visible.value = false
    if(codeRef.value) codeRef.value.setContent()
  }
</script>

<style>
.codeModal .arco-modal-header{
   display: none;
 }
</style>
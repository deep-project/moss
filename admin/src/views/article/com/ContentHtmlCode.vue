<template>
  <a-spin class="w-full h-full" :loading="loading">
    <Codemirror
        v-model="code"
        placeholder=""
        class="w-full"
        :style="{ height: '100%' }"
        :autofocus="false"
        :line-wrapping="true"
        :indent-with-tab="true"
        :autoDestroy="true"
        :tab-size="2"
        :extensions="extensions"
        @ready="onReady"
    />
  </a-spin>
</template>

<script setup>

  import { Codemirror } from 'vue-codemirror'
  import { html } from '@codemirror/lang-html'
  import { oneDark } from '@codemirror/theme-one-dark'
  import { basicSetup } from 'codemirror'

  import prettier from "prettier/standalone";
  import parserHtml from "prettier/parser-html";
  import {useStore} from "@/store";
  import {ref,inject,shallowRef} from 'vue'


  const store = useStore()
  const record = inject('record')
  const code = ref()
  const loading = ref(true)
  const view = shallowRef()
  const extensions = [html(), basicSetup]
  if(store.dark) extensions.push(oneDark)

  function onReady(payload){
    view.value = payload.view
    setTimeout(()=>{
      if(record?.value?.content){
        code.value = prettier.format(record.value.content, { parser: "html", plugins:[parserHtml] })
      }
      loading.value = false
    },600)
  }

  function setContent(){
    if(loading.value === false) record.value.content = code.value
  }

  defineExpose({setContent})

</script>

<style>
.n-spin-content {
  height: 100%;
}
.ͼ1.cm-editor.cm-focused{
  outline: none !important;
}
</style>
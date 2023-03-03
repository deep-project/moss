<template>
  <Toolbar
      ref="toolbarRef"
      class="toolbar"
      :class="{dark:store.dark}"
      :editor="editorRef"
      :defaultConfig="toolbarConfig"
      mode="default"
  />

  <Editor
      class="overflow-y-hidden editor"
      :class="{dark:store.dark}"
      :style="{height: editorHeight}"
      v-model="valueHtml"
      :defaultConfig="editorConfig"
      mode="default"
      @onCreated="handleCreated"
  />
</template>


<script setup>

    import '@wangeditor/editor/dist/css/style.css'
    import {onBeforeUnmount, ref, shallowRef, inject, computed, watch, onMounted} from 'vue'
    import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
    import { i18nChangeLanguage,i18nAddResources } from '@wangeditor/editor'
    import {useElementSize} from '@vueuse/core'
    import {t} from "@/locale";
    import {useStore} from "@/store";
    import {upload} from "@/api";
    import { localeList } from "@/locale"




    const store = useStore()
    const record = inject('record')
    const valueHtml = ref()
    const editorRef = shallowRef()
    const toolbarRef = ref()
    const { height:toolbarHeight } = useElementSize(toolbarRef)
    const editorHeight = computed(()=>'calc(100% - '+(toolbarHeight.value + 2)+'px)')

    const toolbarConfig = {}
    const editorConfig = { placeholder: t('content')+' ......',  MENU_CONF: {} }

    const uploadFun = (file, insertFn)=>{
      let formData = new FormData()
      formData.append("file",file)
      upload(formData).then((resp)=>{
        if(!resp.success || resp.data.length === 0) return
        insertFn(resp.data[0], '', '')
      })
    }

    editorConfig.MENU_CONF['uploadVideo'] = {
      customUpload:uploadFun,
    }
    editorConfig.MENU_CONF['uploadImage'] = {
      customUpload:uploadFun,
    }

    watch(()=>record.value.content, (val)=>{ valueHtml.value = val })
    watch(valueHtml, ()=>{ record.value.content = valueHtml.value })
    watch(()=>store.locale, ()=> setLang())

    onBeforeUnmount(() => {
      const editor = editorRef.value
      if (editor == null) return
      editor.destroy()
    })

    function handleCreated(editor) {
      editorRef.value = editor
      valueHtml.value = record.value.content
    }

    // 添加新语言，如日语 ja
    for(let key in localeList){
      i18nAddResources(key.replace('-','_'), localeList[key].lang.wangEditor)
    }

    setLang()
    function setLang(){
      i18nChangeLanguage(store.locale.replace('-','_')) // 语言key不能使用中横向分割，否则不识别
      // switch (store.locale){
      //   case 'zh-cn':
      //     i18nChangeLanguage('zh-CN')
      //     break;
      //   default:
      //     i18nChangeLanguage('en')
      // }
    }

</script>

<style scoped>
.toolbar,
.editor{
  border: 1px solid var(--color-fill-2);
}
.editor{
  border-top: none;
}
.editor :deep(.w-e-text-container div[data-slate-editor]){
  min-height: 90% !important;
}
.toolbar.dark{
  border: none;
  border-bottom: 3px solid #282c34;
}
.editor.dark{
  border: none;
  border-radius: 3px;
  border-bottom: 1px solid #282c34;
}

.editor.dark :deep(pre>code){text-shadow:0 1px #000}
.editor.dark :deep(pre>code .token.punctuation){color:#999}
.editor.dark :deep(pre>code .token.tag){color: #f85eb4}
.editor.dark :deep(pre>code .token.string){color: #a9f118}
.editor.dark :deep(pre>code .token.url){color: #eaa452}
.editor.dark :deep(pre>code .token.keyword){color: #36baf3}
.editor.dark :deep(pre>code .token.function){color: #f36481}
.editor.dark :deep(pre>code .token.variable){color: #f8b131}
/*
 https://github.com/wangeditor-team/wangEditor/blob/master/packages/editor/src/assets/index.less
*/

.toolbar,
.editor{
  --w-e-toolbar-bg-color:  var(--color-fill-2);
}
.toolbar.dark,
.editor.dark{

  /* textarea - css vars */
  --w-e-textarea-bg-color: var(--color-fill-2);
  --w-e-textarea-color: #9db1c5;

  --w-e-textarea-border-color: #6b7f94;
  --w-e-textarea-slight-color: #d4d4d4;
  --w-e-textarea-slight-bg-color: #2a2f3a;
  --w-e-textarea-slight-border-color: #2a2f3a;
  --w-e-textarea-selected-border-color: #c3ddfd;
  --w-e-textarea-handler-bg-color: #4290f7;

  /* toolbar - css vars */
  --w-e-toolbar-color: #b6c5d4;
  --w-e-toolbar-bg-color: #282c34;
  --w-e-toolbar-border-color: #2d3239;
  --w-e-toolbar-active-color: #fff;
  --w-e-toolbar-active-bg-color: #21252b;
  --w-e-toolbar-disabled-color: #757c83;

  /* modal - css vars */
  --w-e-modal-button-bg-color: #353a44;
  --w-e-modal-button-border-color: #353a44;

}

.editor :deep(h1){
  font-size: 30px;
}
.editor :deep(h2){
  font-size: 24px;
}
.editor :deep(h3){
  font-size: 20px;
}
.editor :deep(h4){
  font-size: 18px;
}
.editor :deep(h5){
  font-size: 16px;
}
</style>
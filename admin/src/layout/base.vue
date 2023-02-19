<template>
  <div class="overflow-hidden fixed" :style="{
    width: store.windowSize.width + 'px',
    height: store.windowSize.height + 'px',
    backgroundColor: store.dark ? store.darkBgColor: store.bgColor
  }">
    <a-config-provider :locale="localeList[store.locale].arcoLang">
      <n-config-provider class="h-full" :locale="localeList[store.locale].naiveLang" :date-locale="localeList[store.locale].naiveDateLang">
        <router-view />
      </n-config-provider>
    </a-config-provider>
    <BgColorPicker />
  </div>
</template>

<script setup>
  import {useRouter} from 'vue-router'
  import {useInitLocale,localeList} from "@/locale";
  import {useStore} from "@/store";
  import {watch} from "vue";
  import BgColorPicker from "@/components/app/BgColorPicker.vue";
  import { NConfigProvider } from 'naive-ui'

  const router = useRouter()
  const store = useStore()

  // 初始化 dark
  if(store.dark) document.body.setAttribute('arco-theme', 'dark');
  // 监听 dark
  watch(()=>store.dark,(val)=>{
    if(val) document.body.setAttribute('arco-theme', 'dark');
    else document.body.removeAttribute('arco-theme');
  })

  // 初始化场景语言
  useInitLocale()

  if(!store.token) router.push({name:"login"})

</script>
<template>
  <div class="absolute bottom-1 right-0 hidden xl:block overflow-hidden text-gray-400 opacity-40 hover:opacity-80 transition">
    <n-color-picker :theme="store.dark ? darkTheme:undefined" class="invisible absolute right-5" v-model:show="show" :value="defaultValue" @update:value="onChange" :swatches="swatches" />
    <icon-palette class="cursor-pointer" :size="24" @click="show = true" />
  </div>
</template>

<script setup>
  import {computed, ref} from "vue";
  import { NColorPicker,darkTheme } from 'naive-ui'
  import {useStore} from "@/store/index.js";

  const store = useStore()
  const show = ref(false)
  const defaultValue = computed(()=>store.dark ? store.darkBgColor:store.bgColor)
  const swatches = computed(()=>store.dark ? ['#21252b','#000000','#100E12']:['#E1E1E1','#F0EBE7','#E7E3F1','#f0f1f6'])

  function onChange(e){
    if(store.dark) store.darkBgColor=e
    else store.bgColor = e
  }


</script>
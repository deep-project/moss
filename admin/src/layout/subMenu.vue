<template>
  <a-layout class="h-full subMenu">
    <a-layout-sider collapsible hide-trigger class="border-r" style="background-color: transparent;border-color: var(--color-border-2)"
                    :width="width"
                    :collapsed="collapsed"
                    :collapsed-width="10"
    >
      <div class="overflow-hidden" :style="{marginRight:collapsed ? '10px':0}">
        <slot name="sider"></slot>
      </div>
      <span class="cursor-pointer absolute z-10 top-1/2 rounded-full flex items-center border p-1 shadow"
            style="right:-12px;margin-top:-22px;color:rgb(var(--arcoblue-5));border-color:var(--color-border-2);background-color:var(--color-bg-2);"
            @click="collapsed=!collapsed">
        <icon-right v-if="collapsed" :size="14" :stroke-width="5" />
        <icon-left v-else :size="14" :stroke-width="5" />
      </span>
    </a-layout-sider>
    <a-layout  style="padding: 0 8px; background-color: var(--color-bg-2)">
      <a-layout-content><slot></slot></a-layout-content>
    </a-layout>
  </a-layout>

</template>


<script setup>
  import {ref} from 'vue'
  import {useStore} from "@/store/index.js";

  const collapsed = ref(false)
  const store = useStore()

  defineProps({width: {type:Number, default:170} })

</script>

<style scoped>
.subMenu :deep(.arco-layout-sider-children){
  overflow: hidden;
}
</style>
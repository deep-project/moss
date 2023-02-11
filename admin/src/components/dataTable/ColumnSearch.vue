<template>

    <div class="p-3 shadow-md" :class="{border:!store.dark}" style="background-color: var(--color-bg-1)">
      <a-space direction="vertical">
        <a-input :model-value="filterValue[0]" :placeholder="$t('search')" @input="(value)=>setFilterValue([value,filterValue[1]])" />
        <div class="flex justify-between select-none">
          <a-checkbox v-model="filterValue[1]">Like</a-checkbox>
          <a-space class="pl-5">
            <a-button size="mini" @click="handleFilterReset">{{ $t('reset') }}</a-button>
            <a-button type="primary" size="mini" @click="handleFilterConfirm">{{ $t('confirm') }}</a-button>
          </a-space>
        </div>
      </a-space>
    </div>

</template>

<script setup>
  import {useStore} from "@/store/index.js";

  defineProps({filterValue:Object,setFilterValue:Function,handleFilterReset:Function,handleFilterConfirm:Function})
  const store = useStore()
</script>
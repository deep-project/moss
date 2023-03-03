<template>

    <div class="w-full flex">
        <a-cascader class="flex-shrink w-full" :options="treeData"
                    v-model="data"
                    :loading="loadingTreeData"
                    :field-names="{value: 'id', label: 'name'}"
                    check-strictly
                    expand-child
                    allow-clear
                    allow-search
                    :placeholder="$t('select')"
                    :style="cascaderStyle" :multiple="multiple"
                @clear="onClear"
        />

      <a-trigger class="flex-grow" trigger="click" position="br" auto-fit-position :unmount-on-close="false" :popup-offset="10">
        <a-button type="primary"><template #icon><icon-edit /></template></a-button>
        <template #content>
          <div class="w-full p-2 border border-gray-600 border-opacity-20" :style="{backgroundColor:'var(--color-bg-5)', maxWidth:'400px'}">
            <a-input-tag v-if="multiple" v-model="data" placeholder="Please Enter ID" allow-clear />
            <a-input-number v-else v-model="data" hide-button allow-clear placeholder="id" @input="onInput" />
          </div>
        </template>
      </a-trigger>

    </div>

</template>


<script setup>

  import {toRefs, ref, watch } from "vue";
  import {useRequest} from "vue-request";
  import {categoryTree} from "@/api/index.js";

  const props = defineProps({ modelValue:Number | Array, cascaderStyle:Object, disabledId:Number, multiple:Boolean })
  const emit = defineEmits(['update:modelValue'])
  const { modelValue } = toRefs(props)
  const data = ref(modelValue.value)

  watch(modelValue,()=>{ data.value = modelValue.value })
  watch(data,(val)=>{ emit('update:modelValue', val)})



  function onInput(val){
    emit('update:modelValue', val)
  }

  function onClear(){
    data.value = 0
  }

  const {data:treeData,loading:loadingTreeData} = useRequest(categoryTree,{
    onSuccess:()=>{
      if(props.disabledId) treeDataDisabled(treeData.value)}
  });


  // treeData禁用自身类目和其子类目
  function treeDataDisabled(list){
    for(let i in list){
      if(list[i].id === props.disabledId){
        list[i].disabled = true
        //treeDataDisabledChildren(list[i].children)
      }else{
        treeDataDisabled(list[i].children)
      }
    }
  }

  // 禁用全部子分类
  // 在属性选择器情况下，子分类依然会显示，使用此方法禁用子分类
  // function treeDataDisabledChildren(list){
  //   if(list){
  //     for(let i in list){
  //       list[i].disabled = true
  //       treeDataDisabledChildren(list[i].children)
  //     }
  //   }
  // }

</script>
<style scoped>
:deep(.arco-select-view-value){
  font-size:12px !important;;
}
</style>
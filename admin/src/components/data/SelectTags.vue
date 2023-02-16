<template>
  <a-trigger position="rt" auto-fit-position :unmount-on-close="false" trigger="click" :popup-translate="[10, 0]">
    <a-input-tag v-model="data" unique-value allow-clear />
    <template #content>
      <a-transfer :data="list" v-model="data" simple  show-search  @search="search" style="background-color: var(--color-bg-5);"/>
    </template>
  </a-trigger>

</template>


<script setup>

  import {toRefs, ref, watch, computed} from "vue";

  import {useRequest} from "vue-request";
  import {tableList, tagGetByIds} from "@/api/index.js";

  const props = defineProps({ modelValue:Array })
  const emit = defineEmits(['update:modelValue'])
  const { modelValue } = toRefs(props)
  const data = ref(modelValue.value)

  watch(modelValue,()=>{ data.value = modelValue.value })
  watch(data,(val)=>{ emit('update:modelValue', val)})


  const { data:selectedData, run:runGetSelected, loading:loadingGetSelected } = useRequest(tagGetByIds, {manual:true})
  const defaultSelected = computed(()=>selectedData.value?.map((item, index) => ({ value: item.id,label: item.name })))

  function getSelected(){
    if(data.value && data.value.length > 0) runGetSelected(data.value)
  }

  const { data:sourceListData, run:runSourceList, loading:loadingSourceList } = useRequest(tableList, {manual:true})
  const sourceList = computed(()=>sourceListData.value?.map((item, index) => ({ value: item.id,label: item.name })))

  // 最终的列表，是由源数据列表和已选择列表组成的
  // 采用转换对象的方式去重合并
  const list = computed(()=>{
    let obj = {}
    let res = []
    for(let k in sourceList.value){
      let item = sourceList.value[k]
      obj[item.value] = item
    }
    for(let k in defaultSelected.value){
      let item = defaultSelected.value[k]
      obj[item.value] = item
    }
    for(let k in obj) res.push(obj[k])
    return res
  })

  function getSourceList(searchValue){
    let params = {page:1, limit:100, order:""}
    if(searchValue) params.where = {field:'name', operator:'like', value:searchValue}
    runSourceList("tag", params)
  }

  getSelected()
  getSourceList()

  function search(value,type){
    if(type === 'source') getSourceList(value)
    getSelected() // 按下搜索时，必须要用选中id再去查一下列表
  }

</script>

<style scoped>
.demo-basic {
  padding: 10px;
  width: 200px;
  background-color: var(--color-bg-popup);
  border-radius: 4px;
  box-shadow: 0 2px 8px 0 rgba(0, 0, 0, 0.15);
}

:deep(.arco-transfer-view){
  height:310px
}
</style>
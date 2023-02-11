<template>

  <a-space wrap fill>
      <a-tag class="tag" v-for="tag of tags" :key="tag.name" closable bordered
             :style="{height:'26px',lineHeight:'24px'}"
             @close="onRemove(tag)">{{tag.name}}</a-tag>

      <a-dropdown v-if="showAddInput" v-model:popupVisible="selectLayerShow" trigger="click" position="bl" :click-to-close="false"
                  @select="onSelect" @popup-visible-change="popupVisibleChange">
        <a-input v-model.trim="addInputValue" class="input" ref="addInputRef" allow-clear
                 :style="{ width: '90px'}" size="mini"
                 @keyup.enter.stop="onCreate" @blur="onInputBlur" @focus="onInputFocus" @input="onInput" />
        <template #content>
            <a-spin class="w-full" style="min-width: 90px" :loading="loadingTagSelectList">
              <div v-if="!tagSelectList || tagSelectList.length === 0">
                <a-empty/>
              </div>
              <template v-else>
                <template v-for="tag in tagSelectList">
                  <a-doption v-if="!existInTags(tag.name)">{{tag.name}}</a-doption>
                </template>
              </template>
            </a-spin>
        </template>
      </a-dropdown>

      <a-button v-else type="outline" size="mini" :style="{ width: '90px', borderStyle:'dashed'}" :disabled="loadingGetTags" @click="onShowAddInput">
        <template #icon><icon-plus /></template>{{$t('create')}}
      </a-button>

  </a-space>

</template>

<script setup>

  import {inject,ref,nextTick,watch } from 'vue'
  import {useRequest} from "vue-request";
  import {
    articleCreateTagByNameList,
    articleDeleteTagByIds, tableList,
    tagsByArticleID
  } from "@/api/index.js";
  import { existName, getIndexByName } from './tag.js'
  import {useDeepCopy} from "@/hooks/utils.js";
  import {listOption} from "@/components/dataTable/index.js";

  const record = inject('record')

  // tags / rawTags
  let rawTags = [] // 原始tags数据
  const {data:tags, run:runGetTags, loading:loadingGetTags} = useRequest(tagsByArticleID, {manual:true, onSuccess:(data)=>{ rawTags = useDeepCopy(data) }})
  const existInTags = (name) => existName(tags.value,name)    // tags中是否已存在某tag
  const existInRawTags = (name) => existName(rawTags, name)   // rawTags中是否已存在某tag
  const pushTagsByName = (name) => { if(!tags.value) tags.value = []; tags.value.push({name:name}) }
  const clearSetTags = (name)=>{ tags.value = tags.value.filter((t) => t.name !== name) }  // 采用重新赋值的方式去掉tags中的项，使页面能够更新
  const findInRawTags = (name) => rawTags[getIndexByName(rawTags, name)]

  // waitCreateTagName / waitRemoveTag
  const waitCreateTagName = ref([]) // 更新时，待创建的标签name
  const waitRemoveTag = ref([])     // 更新时，待移除的标签
  const existInWaitCreateTagName = (name) => waitCreateTagName.value.indexOf(name) > -1
  const existInWaitRemoveTag = (name) => existName(waitRemoveTag.value, name)
  const pushWaitCreateTagName = (name) => waitCreateTagName.value.push(name)
  const pushWaitRemoveTag = (tag) => waitRemoveTag.value.push(tag)
  const removeWaitCreateTagName = (name) => { waitCreateTagName.value.splice(waitCreateTagName.value.indexOf(name), 1)  }
  const removeWaitRemoveTag = (name) =>  { waitRemoveTag.value.splice(getIndexByName(waitRemoveTag.value, name), 1)  }

  function getWaitRemoveTagID(){
    let r = []
    for(let i in waitRemoveTag.value) r.push(waitRemoveTag.value[i].id)
    return r
  }

  // addInput
  const addInputValue = ref('')
  const showAddInput = ref(false)
  const addInputRef = ref(null)
  const inputFocus = ref(false)
  const clearAddInput  = ()=> { addInputValue.value = ""}
  const onShowAddInput = ()=> { showAddInput.value = true ; nextTick(() => addInputRef.value.focus()) }
  const onHideAddInput = ()=> { showAddInput.value = false }
  const onInputFocus   = ()=> { inputFocus.value = true; selectLayerShow.value = true; getTagSelectList() }
  const onInputBlur    = ()=> { inputFocus.value = false; if(!selectLayerShow.value) onCreate() }
  const onInput        = (val)=> { getTagSelectList(val) }


  // create
  function onCreate(){
    let name = addInputValue.value
    clearAddInput()
    onCreateFun(name)
  }

  function onCreateFun(name){
    onHideAddInput()
    if(!name || existInTags(name)) return
    pushTagsByName(name)                                     // 加入tags列表
    if(!record.value.id) return                              // 创建文章，不继续
    if(!existInRawTags(name)) pushWaitCreateTagName(name)    // 不在原始tags列表中，加入待添加列表
    if(existInWaitRemoveTag(name)) removeWaitRemoveTag(name) // 如果在待移除列表中，取消移除
  }

  // remove
  function onRemove(tag){
    let name = tag.name
    clearSetTags(name)          // 去掉tags中的项，使页面能够更新
    if(!record.value.id) return // 创建文章，不继续
    if(!name) return
    if(existInWaitCreateTagName(name)) return removeWaitCreateTagName(name) // 在待创建列表，取消创建
    if(!existInWaitRemoveTag(name)){
      if(tag.id > 0){
        pushWaitRemoveTag(tag)    // 如果存在id，则加入待移除列表
      }else{
        let rawTag = findInRawTags(name) // 根据name查找原始的tag数据
        // 如果存在原始tag，则加入到待移除列表，
        // 否则容易出现：删除原始tag、添加tag、再删除导致无法添加到待移除列表
        if(rawTag) pushWaitRemoveTag(rawTag)
      }
    }
  }


  // callback
  const postRecordGetSuccessCallback = inject('postRecordGetSuccessCallback')
  const createBeforeCallback = inject('createBeforeCallback')
  const updateSuccessCallback = inject('updateSuccessCallback')

  // 如果是更新状态，加载tag列表
  postRecordGetSuccessCallback.value = ()=>{
    if(record.value.id > 0) runGetTags(record.value.id)
  }

  // 创建之前的回调
  createBeforeCallback.value = ()=>{
    record.value.tags = []
    for(let i in tags.value) record.value.tags.push(tags.value[i].name)
  }

  // 更新成功之后的回调
  updateSuccessCallback.value = ()=>{
    if(waitCreateTagName.value.length > 0) createArticleTagByID(record.value.id, waitCreateTagName.value)
    let removeTagIds = getWaitRemoveTagID()
    if(removeTagIds.length > 0) removeArticleTagByID(record.value.id, removeTagIds)
  }

  const {run:createArticleTagByID} = useRequest(articleCreateTagByNameList, {manual:true})
  const {run:removeArticleTagByID} = useRequest(articleDeleteTagByIds, {manual:true})


  // select layer
  const selectLayerShow = ref(false)

  function onSelect(val){
    addInputValue.value = ""
    onCreateFun(val)
  }

  function popupVisibleChange(){
    if(!inputFocus.value) onCreate()
  }

  const { data:tagSelectList, run:runTagSelectList, loading:loadingTagSelectList } = useRequest(tableList, {manual:true})

  function getTagSelectList(searchValue){
    let params = {page:1, limit:100, order:""}
    if(searchValue) params.where = {field:'name', operator:'like', value:searchValue}
    runTagSelectList("tag", params)
  }
</script>

<style scoped>
.input,
 .tag{
   background-color: var(--color-bg-5);
 }
</style>
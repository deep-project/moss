<template>
  <template v-if="selectedKeys.length > 0">
    <a-dropdown :popup-max-height="false" trigger="hover">
      <a-button size="small">
        <template #icon><icon-down /></template>
        {{ $t('action') }}
      </a-button>
      <template #content>

        <a-doption v-if="modelName === 'article'|| modelName === 'category'" @click="setCategoryModelVisible = true">
          <template #icon><icon-edit /></template>
         {{$t('update')}} {{ modelName === 'article' ? $t('category'):$t('parentCategory')}}
        </a-doption>

        <a-doption @click="batchDelete(modelName, selectedKeys)" style="color: rgb(var(--red-6))">
          <template #icon><icon-delete /></template>
          {{$t('delete')}}
        </a-doption>
      </template>
    </a-dropdown>
    <a-tag class="ml-4 overflow-hidden" style="border-radius:18px;background-color: var(--color-fill-2)"
           @close="selectedKeys = []" closable>{{$t('checked')}} {{ selectedKeys.length }}</a-tag>
  </template>

  <a-modal v-if="modelName === 'article'|| modelName === 'category'" v-model:visible="setCategoryModelVisible" @ok="setCategoryHandle" simple>
    <SelectCategory v-model="setCategoryValue" />
  </a-modal>

</template>


<script setup>
  import {inject,ref} from 'vue'
  import {useRequest} from "vue-request";
  import {articleBatchSetCategory, categoryBatchSetParentCategory} from "@/api/index.js";
  import {Message} from '@arco-design/web-vue'
  import {t} from "@/locale/index.js";
  import SelectCategory from "@/components/data/SelectCategory.vue"


  const selectedKeys = inject('selectedKeys')
  const batchDelete = inject('batchDelete')
  const modelName = inject('modelName')
  const refresh = inject('refresh')

  const setCategoryModelVisible = ref(false)
  const setCategoryValue = ref()

  function setCategoryHandle(){
    if(modelName === 'article')  setArticlesCategory(setCategoryValue.value, selectedKeys.value)
    if(modelName === 'category')  setCategoryParentCategory(setCategoryValue.value, selectedKeys.value)
  }

  let opt = { manual:true,
    onSuccess:(resp)=>{
      if(!resp.success) return
      selectedKeys.value = []
      refresh()
      Message.success(t('message.success',[t('update')]))
    }}
  const { run:setArticlesCategory } = useRequest(articleBatchSetCategory, opt)
  const { run:setCategoryParentCategory } = useRequest(categoryBatchSetParentCategory, opt)


</script>
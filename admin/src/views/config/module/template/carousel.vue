<template>

  <a-form-item :label="$t('enable')">
    <a-switch type="round" v-model="data.enable_carousel" />
  </a-form-item>

  <a-grid v-if="data.enable_carousel" :cols="{ xs: 1, sm: 1, md: 1, lg: 2, xl: 3, xxl: 4 }" :colGap="12" :rowGap="16">
    <a-grid-item v-for="(item,k) in data.carousel">
      <a-card :title="'#'+(k+1)" size="small" style="height: 324px;" :header-style="{padding:'8px'}">
        <template #extra>
            <a-button-group size="small">
              <a-button :disabled="k===0" @click="up(k)">
                <template #icon><icon-caret-left /></template>
              </a-button>
              <a-button :disabled="k===data.carousel.length -1" @click="down(k)">
                <template #icon><icon-caret-right /></template>
              </a-button>
            </a-button-group>
            <a-button @click="remove(k)" class="ml-2" size="small">
              <template #icon><icon-close-circle /></template>
            </a-button>
        </template>
        <a-form :model="item" layout="vertical">
          <a-form-item field="image" :label="$t('image')">
            <UploadImgInput v-model="item.image" class="w-full" />
          </a-form-item>
          <a-form-item field="link" :label="$t('link')">
            <a-input v-model="item.link" />
          </a-form-item>
          <a-form-item field="title" :label="$t('title')">
            <a-input v-model="item.title" />
          </a-form-item>
        </a-form>
      </a-card>
    </a-grid-item>
    <a-grid-item>
      <a-card style="height: 324px;background-color: var(--color-fill-1);" class="h-full w-full flex items-center justify-center cursor-pointer" @click="add">
        <div class="text-gray-300"><icon-plus-circle :size="60" /></div>
      </a-card>
    </a-grid-item>
  </a-grid>
</template>

<script setup>
  import {inject} from 'vue'
  import UploadImgInput from '@/components/utils/UploadImgInput.vue'

  const data = inject('data')
  if(!data.value.carousel) data.value.carousel= []

  function add(){
    data.value.carousel.push({image:'',link:'',title:''})
  }

  function remove(k){
    data.value.carousel.splice(k, 1)
  }

  const swapItems = function(arr, index1, index2){
    arr[index1] = arr.splice(index2,1,arr[index1])[0]
    return arr
  }

  function up(index) {
    let newArr =[]
    if (data.value.carousel.length > 1 && index !== 0) {
      newArr = swapItems(data.value.carousel, index, index - 1)
    }
    if(newArr.length > 0) data.value.carousel = newArr
  }

  function down(index) {
    let newArr =[]
    if (data.value.carousel.length > 1 && index !== (data.value.carousel.length - 1)) {
      newArr = swapItems(data.value.carousel, index, index + 1)
    }
    if(newArr.length > 0) data.value.carousel = newArr
  }

</script>
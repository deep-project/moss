<template>
  <a-input-group>
    <a-input :style="inputStyle" v-model="img" allow-clear>
      <template #suffix>
        <a-link @click="fileInput.click()" :hoverable="true"><icon-upload /></a-link>
      </template>
    </a-input>
    <input class="hidden" type="file" ref="fileInput" accept="image/*" @change="getFile">
  </a-input-group>
</template>

<script setup>
  import {ref, toRefs, watch} from 'vue'
  import {useRequest} from "vue-request";
  import {upload as uploadAPI} from "@/api/index.js";

  const props = defineProps({ modelValue:String, inputStyle:String })
  const {modelValue} = toRefs(props)
  const emit = defineEmits(['update:modelValue'])
  const img = ref(modelValue.value)
  const fileInput= ref()

  watch(modelValue,()=>{ img.value = modelValue.value })
  watch(img,()=>{
    emit('update:modelValue', img.value)
  })

  const {run:upload} = useRequest(uploadAPI,{
    manual:true,
    onBefore:()=>{},
    onSuccess:(resp)=>{
      if(resp.success && resp.data.length > 0) {
        img.value = resp.data[0]
      }
    }})



  function getFile(event) {
    let formFile = new FormData();
    for(let f of event.target.files){
      formFile.append("file", f)
    }
    upload(formFile)
    fileInput.value.value = ""
  }

</script>
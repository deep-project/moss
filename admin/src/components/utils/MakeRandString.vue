<template>
  <icon-loop @click="onMake" class="cursor-pointer" />
</template>

<script setup>
  import {useRandString} from "@/hooks/utils";

  const emit = defineEmits(['update:modelValue'])
  const props = defineProps({
    modelValue:String,
    title:  { type: String, default: ''},
    length: { type: Number, default: 10},
    chars:  { type: String, default: 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz'}})

  const onMake = ()=>{
    emit('update:modelValue', useRandString(props.length, props.chars))
  }

</script>
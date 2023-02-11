<template>
    <a-form-item class="my-5" :label="label ? label :$t('select')">
      <a-select class="w-64" v-model="value.active" :options="range" />
    </a-form-item>
    <a-card class="mb-5 max-w-2xl py-5">
      <div v-for="(obj,key) in value.driver">
        <component v-if="key === value.active" v-bind:is="getDriver(key)" :data="obj"></component>
      </div>
    </a-card>
</template>

<script setup>
  import {defineAsyncComponent} from 'vue'

  defineProps({value:Object,range:Array,label:String})

  const drivers = import.meta.glob("./drivers/*.vue");

  function getDriver(name){
    let d = drivers["./drivers/"+name+".vue"]
    if(!d) return
    return defineAsyncComponent(d)
  }
</script>
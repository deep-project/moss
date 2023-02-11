<template>
  <a-form-item class="my-5" :label="$t('storage')">
    <a-select class="w-64" v-model="data.active_driver" :options="['badger','redis','memcached']" />
  </a-form-item>
  <a-card class="mb-5 max-w-2xl py-5">
    <div v-for="(obj,key) in data.driver">
      <div v-if="key === data.active_driver" class="overflow-auto" style="max-height: 280px">
        <component v-bind:is="getDriver(key)" :data="obj"></component>
      </div>
    </div>
  </a-card>
</template>

<script setup>
  import {defineAsyncComponent} from 'vue'
  import {inject} from 'vue'

  const data = inject('data')
  const drivers = import.meta.glob("./drivers/*.vue");

  function getDriver(name){
    let d = drivers["./drivers/"+name+".vue"]
    if(!d) return
    return defineAsyncComponent(d)
  }
</script>
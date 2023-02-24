<template>

  <a-form-item label="secret">
    <a-input class="w-64" v-model="data.prof_secret" >
      <template #suffix><MakeRandString v-model="data.prof_secret" :length="10" /></template>
    </a-input>
    <a-tag v-if="data.prof_secret===''" color="red" class="ml-1"><template #icon><icon-exclamation-circle-fill /></template>unsafe</a-tag>
  </a-form-item>

  <a-descriptions title="Usage" bordered :column="1" size="large">
    <a-descriptions-item label="web">
      <span @click="useOpenLink" class="inline-block cursor-pointer hover:underline underline-offset-4 decoration-2 hover:text-blue-500">
        {{domain}}{{path}}/debug/pprof
      </span>
    </a-descriptions-item>
    <a-descriptions-item label="heap">
      <a-tag @click="useCopy" class="cursor-pointer">go tool pprof -http=:9090 {{domain}}{{path}}/debug/pprof/heap</a-tag>
    </a-descriptions-item>
    <a-descriptions-item label="goroutine">
      <a-tag @click="useCopy" class="cursor-pointer">go tool pprof -http=:9090 {{domain}}{{path}}/debug/pprof/goroutine</a-tag>
    </a-descriptions-item>

  </a-descriptions>

</template>


<script setup>

  import {computed, inject} from "vue";
  import MakeRandString from "@/components/utils/MakeRandString.vue";
  import {useOpenLink,useCopy} from '@/hooks/utils.js'
  import {useSiteURL} from '@/hooks/app/index.js'
  import {useStore} from '@/store'

  const data = inject('data')
  const store = useStore()
  const domain = useSiteURL(store)

  const path = computed(()=> data.value.prof_secret ? "/" + data.value.prof_secret :"")

</script>
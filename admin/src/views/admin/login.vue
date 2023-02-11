<template>
  <div class="w-screen h-screen">
    <Login :show-lang-select="true" @success="onSuccess" />
  </div>
</template>

<script setup>
  import Login from '@/components/admin/Login.vue'
  import {useRouter} from "vue-router";

  const router = useRouter()
  function onSuccess(){
    router.push({name:"index"})
  }
</script>
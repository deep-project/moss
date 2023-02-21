<template>
  <a-space direction="vertical" size="medium" fill>
    <div class="w-full rounded-lg overflow-hidden"><System /></div>
    <div class="w-full rounded-lg overflow-hidden"><App /></div>
    <div class="w-full rounded-lg overflow-hidden"> <Article /></div>
    <div class="w-full rounded-lg overflow-hidden"> <Other /></div>
    <div class="w-full rounded-lg overflow-hidden"> <About /></div>
  </a-space>
</template>

<script setup>
  import System from './system.vue'
  import App from './app.vue'
  import Article from './article.vue'
  import Other from './other.vue'
  import About from './about.vue'

</script>

<style>
.arco-card-header{
  border:none;
}
.arco-statistic-title{
  margin-bottom:0;
  color:#6b6b6b;
}
.arco-card-size-medium .arco-card-header{
  padding-top:16px;
}
.arco-card-size-medium .arco-card-body{
  padding-top:26px;
  padding-bottom:26px;
}
</style>
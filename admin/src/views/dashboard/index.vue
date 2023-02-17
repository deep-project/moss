<template>
  <a-space direction="vertical" size="medium" fill>

    <a-grid :cols="{ xs: 1, sm: 1, md: 1, lg:2 }" :colGap="12" :rowGap="12">
      <a-grid-item><System /></a-grid-item>
      <a-grid-item><App /></a-grid-item>
    </a-grid>

    <Article />
    <Other />
    <About />
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
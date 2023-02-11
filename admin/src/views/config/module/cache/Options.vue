<template>
  <a-form-item :label="$t('home')" class="mt-5">
    <Item :data="data.options.home" name="home" />
  </a-form-item>

  <a-form-item :label="$t('article')">
    <Item :data="data.options.article" name="article" />
  </a-form-item>

  <a-form-item :label="$t('category')">
    <Item :data="data.options.category" name="category" />
  </a-form-item>

  <a-form-item :label="$t('tag')">
    <Item :data="data.options.tag" name="tag" />
  </a-form-item>

  <a-form-item :label="$t('sitemap')">
    <Item :data="data.options.sitemap" name="sitemap" />
  </a-form-item>

  <a-form-item :label="$t('templatePage')">
    <Item :data="data.options.page" name="page" />
  </a-form-item>
</template>

<script setup>
  import Item from "./Item.vue";
  import {inject} from 'vue';

  const data = inject('data')
</script>
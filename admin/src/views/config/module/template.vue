<template>
  <a-tabs type="card-gutter" lazy-load>
    <a-tab-pane key="content" :title="$t('content')">
      <div class="p-3"><Content /></div>
    </a-tab-pane>

    <a-tab-pane key="menu" :title="$t('menu')">
      <div class="p-3"><Menu /></div>
    </a-tab-pane>

    <a-tab-pane key="carousel" :title="$t('carousel')">
      <div class="p-3"><Carousel /></div>
    </a-tab-pane>

    <a-tab-pane key="indexList" :title="$t('homePageList')">
      <div class="p-3"><List :data="data.index_list" /></div>
    </a-tab-pane>

    <a-tab-pane key="globalList" :title="$t('globalList')">
      <div class="p-3"><List :data="data.global_list" /></div>
    </a-tab-pane>

    <a-tab-pane key="categoryPageList" :title="$t('categoryPageList')">
      <div class="p-3"><List :data="data.category_page_list" :showPageOption="true" /></div>
    </a-tab-pane>

    <a-tab-pane key="tagPageList" :title="$t('tagPageList')">
      <div class="p-3"><List :data="data.tag_page_list" :showPageOption="true" :disableOrder="true" /></div>
    </a-tab-pane>

    <a-tab-pane key="tagCloud" :title="$t('tagCloud')">
      <div class="p-3"><TagCloud /></div>
    </a-tab-pane>

  </a-tabs>
</template>


<script setup>
  import Content from './template/content.vue'
  import Carousel from './template/carousel.vue'
  import Menu from './template/menu.vue'
  import List from './template/List.vue'
  import TagCloud from './template/tagCloud.vue'
  import {inject} from "vue";

  const data = inject('data')


</script>
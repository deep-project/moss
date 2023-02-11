<template>

  <a-tabs type="card-gutter" animation class="mb-4" style="max-width: 600px">
    <a-tab-pane key="article" :title="$t('article')" class="p-5">
      <Option :data="data.article" />
    </a-tab-pane>
    <a-tab-pane key="category" :title="$t('category')" class="p-5">
      <Option :data="data.category" />
    </a-tab-pane>
    <a-tab-pane key="tag" :title="$t('tag')" class="p-5">
      <Option :data="data.tag" />
    </a-tab-pane>
    <a-tab-pane key="router" :title="$t('router')" class="p-5">
      <SitemapURL class="w-full" />
    </a-tab-pane>
  </a-tabs>

</template>

<script setup>
  import Option from './sitemap/option.vue'
  import SitemapURL from './sitemap/url.vue'
  import {inject} from 'vue'

  const data = inject('data')
</script>
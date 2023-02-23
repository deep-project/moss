<template>

  <a-form-item label="Article Views Pool" help="update article views pool capacity">
    <a-input-number class="w-64" v-model="data.article_views_pool" />
  </a-form-item>

  <a-form-item label="Fast Offset" help="fast offset min page">
    <a-input-number class="w-64" v-model="data.fast_offset_min_page" />
  </a-form-item>

  <a-form-item label="Categories Limit" help="view all categories limit">
    <a-input-number class="w-64" v-model="data.view_all_category_limit" />
  </a-form-item>

  <a-form-item label="Unique Title" help="article title must be unique">
    <a-switch type="round" v-model="data.unique_title" />
  </a-form-item>

  <a-divider />

</template>

<script setup>

  import {inject} from 'vue'

  const data = inject('data')

</script>
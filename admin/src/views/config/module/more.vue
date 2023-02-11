<template>

  <a-form-item label="Article Views Pool" help="update article views pool capacity">
    <a-input-number class="w-64" v-model="data.article_views_pool" />
  </a-form-item>

  <a-form-item label="Fast Offset" help="fast offset min page">
    <a-input-number class="w-64" v-model="data.fast_offset_min_page" />
  </a-form-item>

  <a-form-item label="Categories limit" help="view all categories limit">
    <a-input-number class="w-64" v-model="data.view_all_category_limit" />
  </a-form-item>


</template>

<script setup>

  import {inject} from 'vue'

  const data = inject('data')

</script>
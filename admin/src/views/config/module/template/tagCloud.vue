<template>
  <div class="w-full" style="max-width: 360px">

    <a-form-item :label="$t('limit')">
      <a-input-number v-model="data.tag_cloud.limit" :min="0"></a-input-number>
    </a-form-item>

    <a-form-item :label="$t('order')">
      <a-select v-model="data.tag_cloud.order">
        <a-option value="">none</a-option>
        <a-option value="id desc">id desc</a-option>
        <a-option value="id asc">id asc</a-option>
      </a-select>
    </a-form-item>
    <a-form-item :label="$t('select')">
      <SelectTags v-model="data.tag_cloud.select"  />
    </a-form-item>

  </div>
</template>

<script setup>
  import {inject} from "vue";
  import SelectTags from '@/components/data/SelectTags.vue'

  const data = inject('data')

</script>
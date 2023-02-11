<template>
  <a-form-item label="预生成文章缓存">
    <a-switch type="round" v-model="data.pre_build_article" />
    <a-tooltip content="开启后，编辑或添加文章会主动生成缓存">
      <span class="cursor-pointer text-gray-500 ml-2"><icon-question-circle /></span>
    </a-tooltip>
  </a-form-item>
</template>

<script setup>
  import {inject} from 'vue'

  const data = inject('data')
</script>
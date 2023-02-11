<template>
  <a-form-item field="title" hide-label :rules="[{ required:true, message: $t('message.required',[$t('title')]) }]">
    <a-input v-model="record.title" :placeholder="$t('title')+'...'" :style="{height:'46px'}" :input-attrs="{style:'font-size:20px'}" :max-length="250" allow-clear show-word-limit />
  </a-form-item>
  <div class="overflow-hidden relative z-50" style="height: calc(100% - 66px)">
    <Content />
  </div>
</template>

<script setup>
  import {inject} from "vue";
  import Content from "@/views/article/com/Content.vue";
  const record = inject('record')
</script>
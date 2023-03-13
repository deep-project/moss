<template>
  <a-descriptions :column="1" bordered>
    <a-descriptions-item label="time"><n-time :time="data.time" /></a-descriptions-item>
    <a-descriptions-item label="level">{{data.level}}</a-descriptions-item>
    <a-descriptions-item label="file">{{data.file}}</a-descriptions-item>
    <a-descriptions-item label="msg">{{data.msg}}</a-descriptions-item>
    <template v-for="(val,key) in data.more">
      <a-descriptions-item v-if="key!=='buf'" :label="key">{{val}}</a-descriptions-item>
    </template>
    <a-descriptions-item v-if="data.buf" label="buf">
      <code class="font-mono text-gray-500 whitespace-pre-line">{{data.buf}}</code>
    </a-descriptions-item>

  </a-descriptions>
</template>

<script setup>
  import { NTime } from 'naive-ui'

  defineProps({data:Object})

</script>
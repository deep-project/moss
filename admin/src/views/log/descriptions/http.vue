<template>
  <a-descriptions :column="1" bordered>
    <a-descriptions-item label="time"><n-time :time="data.time" /></a-descriptions-item>
    <a-descriptions-item label="take">{{data.take}} ms</a-descriptions-item>
    <a-descriptions-item label="status">{{data.status}}</a-descriptions-item>
    <a-descriptions-item label="depth">{{data.depth}}</a-descriptions-item>
    <a-descriptions-item label="method">{{data.method}}</a-descriptions-item>
    <a-descriptions-item label="ip">{{data.ip}}</a-descriptions-item>
    <a-descriptions-item label="region">{{data.region}}</a-descriptions-item>
    <a-descriptions-item label="url">
      <a class="hover:underline cursor-pointer hover:text-blue-400" @click="useOpenLink(data.url)">{{data.url}}</a>
    </a-descriptions-item>
    <a-descriptions-item label="referer">
      <a class="hover:underline cursor-pointer hover:text-blue-400" @click="useOpenLink(data.referer)">{{data.referer}}</a>
    </a-descriptions-item>
    <a-descriptions-item label="user-agent">{{data.userAgent}}</a-descriptions-item>
    <a-descriptions-item label="headers">
      <code class="font-mono text-gray-500 whitespace-pre">{{data.headers}}</code>
    </a-descriptions-item>
  </a-descriptions>
</template>

<script setup>
  import { NTime } from 'naive-ui'
  import {useOpenLink} from "@/hooks/utils.js";

  defineProps({data:Object})

</script>
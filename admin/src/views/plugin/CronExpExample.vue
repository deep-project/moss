<template>
  <a-descriptions bordered :column="1">
    <a-descriptions-item v-for="item in cronExpExample" :label="item.explain">
      <a-space><a-tag class="cursor-pointer" v-for="val in item.example" @click="useCopy">{{val}}</a-tag></a-space>
    </a-descriptions-item>
  </a-descriptions>
  <a-alert banner class="mt-5">
    <span class="text-xs">更多设置方法,参考cron表达式,注意: cron表达式只支持6位(不支持年)</span>
  </a-alert>
</template>

<script setup>
  import {useCopy} from "@/hooks/utils.js";


  const cronExpExample = [
    {explain:'每30秒', example:['@every 30s','0/30 * * * * ? ']},
    {explain:'每5分钟', example:['@every 5m']},
    {explain:'每1个半小时', example:['@every 1h30m']},
    {explain:'每小时', example:['@hourly']},
    {explain:'每天', example:['@daily','@midnight']},
    {explain:'每天3:15分', example:['0 15 3 * * ?']},
    {explain:'每周', example:['@weekly']},
    {explain:'每周二 0:30分', example:['0 30 0 * * 2']},
    {explain:'每周四上午7点', example:[' 0 0 7 ? * 4']},
    {explain:'每月', example:['@monthly']},
    {explain:'每年', example:['@yearly', '@annually']},
  ]

</script>
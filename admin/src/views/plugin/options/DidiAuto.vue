<template>

  <a-form-item label="站点id" required>
    <a-input-number v-model="data.site_id" class="input" :min="0" />
  </a-form-item>

  <a-form-item label="cookie" required>
    <a-textarea v-model="data.cookie" />
  </a-form-item>

  <a-divider />

  <a-form-item label="延迟检测时间">
    <a-input-number v-model="data.detect_delay" class="input" :min="0" /> <span class="text-sm text-gray-400 ml-3">{{$t('minutes')}}</span>
  </a-form-item>

  <a-form-item :label="$t('retry')">
    <a-input-number v-model="data.retry" class="input" :min="0" />
  </a-form-item>

  <a-form-item label="重试等待">
    <a-input-number v-model="data.retry_sleep" class="input" :min="0" /> <span class="text-sm text-gray-400 ml-3">{{$t('seconds')}}</span>
  </a-form-item>

  <a-form-item label="清除缓存URL" help="清除首页缓存触发的URL">
    <a-input v-model="data.clear_cache_url" class="input" />
  </a-form-item>

  <a-form-item label="清除缓存后等待" v-if="data.clear_cache_url">
    <a-input-number v-model="data.clear_cache_sleep" class="input" :min="0" /> <span class="text-sm text-gray-400 ml-3">{{$t('seconds')}}</span>
  </a-form-item>

</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")
</script>
<style scoped>
.input{
  width: 220px;
}
</style>


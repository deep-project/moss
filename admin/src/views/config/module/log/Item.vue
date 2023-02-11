<template>
  <div class="p-3">
      <a-form-item v-if="data.id !=='plugin'" :label="$t('enable')">
        <a-switch type="round" v-model="data.enable" />
      </a-form-item>

      <template v-if="data.enable">
        <a-form-item :label="$t('maxSize')">
          <a-input-number class="w-64" v-model="data.max_size" /><span class="ml-2 text-gray-400">MB</span>
        </a-form-item>
        <a-form-item :label="$t('maxAge')">
          <a-input-number class="w-64" v-model="data.max_age" /><span class="ml-2 text-gray-400">Day</span>
        </a-form-item>
        <a-form-item :label="$t('maxBackups')">
          <a-input-number class="w-64" v-model="data.max_backups" />
        </a-form-item>
    <!--    <a-form-item label="file path">-->
    <!--      <a-input class="w-64" disabled v-model="data.file_path" />-->
    <!--    </a-form-item>-->
        <a-form-item :label="$t('compress')">
          <a-switch  type="round" v-model="data.compress" />
        </a-form-item>

        <a-form-item v-if="data.id ==='app'" :label="$t('level')">
          <div class="w-64"><a-select v-model="data.level" :options="levelOptions" /></div>
        </a-form-item>
        <a-form-item v-if="data.id ==='slow_sql'" :label="$t('slowSqlThreshold')">
          <a-input-number class="w-64" v-model="gdata.slow_sql_threshold" />
        </a-form-item>
        <a-form-item v-if="data.id ==='spider'" :label="$t('spiderFeature')">
          <a-input-tag  v-model="gdata.spider_feature" />
        </a-form-item>
      </template>
      <a-alert v-if="data.id ==='plugin'">需重启应用后生效</a-alert>
  </div>
</template>


<script setup>

  import {inject} from 'vue'

  defineProps({data:Object})
  const gdata = inject('data')

  const levelOptions = [
    {label:"debug",value:"debug"},
    {label:"info",value:"info"},
    {label:"warn",value:"warn"},
    {label:"error",value:"error"},
    {label:"panic",value:"panic"},
    {label:"fatal",value:"fatal"},
  ]
</script>
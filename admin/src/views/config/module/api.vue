<template>
  <a-tabs @change="tabsHandle">
    <a-tab-pane key="options" :title="$t('options')">
      <a-form-item :label="$t('enable')">
        <a-switch type="round" v-model="data.enable" />
      </a-form-item>
      <template v-if="data.enable">
          <a-form-item label="token">
            <a-input v-model="data.secret_key" class="w-64">
              <template #suffix><MakeRandString v-model="data.secret_key" :length="20" /></template>
            </a-input>
          </a-form-item>
          <a-alert :title="$t('warning')" type="warning" class="mb-5">{{$t('apiWarning')}}</a-alert>
          <a-alert :title="$t('auth')" type="info" class="mb-5">
            <a-descriptions  :column="1">
              <a-descriptions-item label="By Request Header">
                <a-tag bordered><code>token: {{data.secret_key}}</code></a-tag>
              </a-descriptions-item>
              <a-descriptions-item label="By Request URL">
                <a-tag bordered><code>/xxx/post?token={{data.secret_key}}</code></a-tag>
                <a-tag color="red" class="ml-1"><template #icon><icon-exclamation-circle-fill /></template>unsafe</a-tag>
              </a-descriptions-item>
            </a-descriptions>
          </a-alert>
      </template>
    </a-tab-pane>

   <a-tab-pane key="documentation" :title="$t('documentation')">
     <Documentation class="mb-5" />
    </a-tab-pane>
  </a-tabs>

</template>

<script setup>
  import MakeRandString from "@/components/utils/MakeRandString.vue";
  import Documentation from "@/views/config/module/api/Documentation.vue";
  import {inject} from 'vue'

  const data = inject('data')
  const showBtn = inject('showBtn')

  function tabsHandle(val){
    showBtn.value = val !== "documentation";
  }

</script>
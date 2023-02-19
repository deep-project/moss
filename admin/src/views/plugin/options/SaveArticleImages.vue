<template>
  <a-form-item :label="$t('enable')">
    <a-space>
      <a-switch v-model="data.enable_on_create" type="round">
        <template #checked>{{ $t('onCreate') }}</template>
        <template #unchecked>{{ $t('onCreate') }}</template>
      </a-switch>
      <a-switch v-model="data.enable_on_update" type="round">
        <template #checked>{{ $t('onUpdate') }}</template>
        <template #unchecked>{{ $t('onUpdate') }}</template>
      </a-switch>
    </a-space>
  </a-form-item>

  <template v-if="data.enable_on_create || data.enable_on_update">

  <a-divider class="w-full" style="margin-top:0" />

  <a-tabs type="rounded">
    <a-tab-pane key="base" title="Base">
      <a-form-item label="max width">
        <a-input-number v-model="data.max_width" class="numberInput" :min="0" />
      </a-form-item>
      <a-form-item label="max height">
        <a-input-number v-model="data.max_height" class="numberInput" :min="0" />
      </a-form-item>
      <a-form-item label="thumb width">
        <a-input-number v-model="data.thumb_width" class="numberInput" :min="0" />
      </a-form-item>
      <a-form-item label="thumb height">
        <a-input-number v-model="data.thumb_height" class="numberInput" :min="0" />
      </a-form-item>
      <a-form-item label="thumb min width">
        <a-input-number v-model="data.thumb_min_width" class="numberInput" :min="0" />
      </a-form-item>
      <a-form-item label="thumb min height">
        <a-input-number v-model="data.thumb_min_height" class="numberInput" :min="0" />
      </a-form-item>
    </a-tab-pane>
    <a-tab-pane key="more" title="More">
      <a-form-item label="down retry">
        <a-input-number v-model="data.down_retry" class="numberInput" :min="0" :max="10" />
      </a-form-item>
      <a-form-item label="always resize">
        <a-switch type="round" v-model="data.always_resize"/>
      </a-form-item>
      <a-form-item label="thumb extract focus">
        <a-switch type="round" v-model="data.thumb_extract_focus"/>
      </a-form-item>
      <a-form-item label="remove down fail">
        <a-switch type="round" v-model="data.remove_if_down_fail"/>
      </a-form-item>

      <a-form-item label="down proxy">
        <a-input v-model="data.down_proxy" class="w-full" />
      </a-form-item>

      <a-form-item label="down referer">
        <a-textarea v-model="data.down_referer" :auto-size="{minRows:4,maxRows:6}"/>
      </a-form-item>
    </a-tab-pane>
  </a-tabs>

  </template>

</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")

</script>

<style scoped>
.numberInput{
  width: 220px;
}
</style>

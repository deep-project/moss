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

  <a-divider class="w-full" style="margin-top:0" />

  <a-form-item label="allow relative urls">
    <a-switch v-model="data.allow_relative_urls" type="round" />
  </a-form-item>

  <a-form-item label="links nofollow" help="require no follow on links">
    <a-switch v-model="data.require_no_follow_on_links" type="round" />
  </a-form-item>


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

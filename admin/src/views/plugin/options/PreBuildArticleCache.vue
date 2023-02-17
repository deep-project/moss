<template>
  <a-form-item :label="$t('enable')">
    <a-space>
      <a-switch v-model="data.enable_on_create" type="round">
        <template #checked>{{ $t('create') }}</template>
        <template #unchecked>{{ $t('create') }}</template>
      </a-switch>
      <a-switch v-model="data.enable_on_update" type="round">
        <template #checked>{{ $t('update') }}</template>
        <template #unchecked>{{ $t('update') }}</template>
      </a-switch>
    </a-space>
  </a-form-item>

  <a-divider class="w-full" style="margin-top:0" />


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

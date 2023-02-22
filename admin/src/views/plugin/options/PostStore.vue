<template>


  <a-form-item :label="$t('limit')">
    <a-input-number v-model="data.limit" :style="{ width: '250px' }" :min="0" />
  </a-form-item>

  <a-form-item :label="$t('order')">
    <a-radio-group v-model="data.order">
      <a-radio :value="0">{{ $t('new') }}</a-radio>
      <a-radio :value="1">{{ $t('old') }}</a-radio>
      <a-radio :value="2">{{ $t('random') }}</a-radio>
    </a-radio-group>
  </a-form-item>

  <a-form-item :label="$t('category')">
    <SelectCategory v-model="data.category_ids" multiple  :style="{ width: '282px' }" />
  </a-form-item>

  <a-form-item :label="$t('delete')" help="delete on failure">
      <a-switch v-model="data.delete_on_failure" type="round" />
  </a-form-item>

</template>


<script setup>
 import {inject} from "vue";
 import SelectCategory from "@/components/data/SelectCategory.vue"

 const data = inject("options")
</script>

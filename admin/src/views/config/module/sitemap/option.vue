<template>

  <a-form-item :label="$t('limit')">
    <a-input-number v-model="data.limit" class="w-64" :min="0" />
  </a-form-item>

  <a-form-item :label="$t('changeFreq')">
    <a-select v-model="data.change_freq" :options="changeFreqOption" class="w-64" />
  </a-form-item>

  <a-form-item :label="$t('priority')">
    <a-input-number v-model="data.priority" class="w-64" :min="0" :max="1.0" :step="0.1" :precision="1" />
  </a-form-item>

  <a-form-item :label="$t('withinHours')">
    <a-input-number v-model="data.in_hours" class="w-64" :min="0" placeholder="all" />
  </a-form-item>


</template>


<script setup>

  import {useStore} from "@/store";
  import {computed} from "vue";

  defineProps({data:Object})
  const store = useStore()

  const changeFreqOption = computed(()=>[
    { label: "always", value: 'always' },
    { label: "hourly", value: 'hourly' },
    { label: "daily", value: 'daily' },
    { label: "weekly", value: 'weekly' },
    { label: "monthly", value: 'monthly' },
    { label: "yearly", value: 'yearly' },
    { label: "never", value: 'never' },
  ])


</script>

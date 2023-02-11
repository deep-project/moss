<template>
  <a-input-group>
    <a-input-number style="width: 150px" v-model="data.number" />
    <a-select v-model="data.unit" :options="units" :style="{width:'106px'}"  />
  </a-input-group>
</template>

<script setup>
  import {t} from '@/locale'
  import {computed} from "vue";

  const {data} = defineProps({data:Object})
  const units = computed(()=>[
    {label: t("seconds"), value: 'second'},
    {label: t("minutes"), value: 'minute'},
    {label: t("hours"), value: 'hour'},
    {label: t("days"), value: 'day'},
  ])

</script>

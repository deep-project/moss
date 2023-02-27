<template>

    <a-form-item label="path">
     <a-input class="w-64" v-model="data.path" />
    </a-form-item>

    <a-form-item label="valueLogLoadingMode" help="default: MemoryMap" >
      <a-radio-group v-model="data.valueLogLoadingMode" :options="loadingModeOptions" />
    </a-form-item>

    <a-form-item label="tableLoadingMode" help="default: MemoryMap" >
      <a-radio-group v-model="data.tableLoadingMode" :options="loadingModeOptions"/>
    </a-form-item>

    <a-form-item label="numMemtables" help="default: 2">
      <a-input-number class="w-64" v-model="data.numMemtables" :min="1"></a-input-number>
    </a-form-item>

    <a-form-item label="maxTableSize" help="default: 16">
      <a-input-number class="w-64" v-model="data.maxTableSize" :min="1">
        <template #suffix>M</template>
      </a-input-number>
      <a-tag class="ml-2" color="blue">RAM: {{data.numMemtables}} * {{data.maxTableSize}}M ≈ {{data.numMemtables * data.maxTableSize}}M</a-tag>
    </a-form-item>

    <a-form-item label="valueLogFileSize" help="default: 256">
      <a-input-number class="w-64" v-model="data.valueLogFileSize" :min="0">
        <template #suffix>M</template>
      </a-input-number>
    </a-form-item>


    <a-form-item label="numCompactors" help="default: 2">
      <a-input-number class="w-64" v-model="data.numCompactors" :min="2"></a-input-number>
    </a-form-item>

  <a-form-item label="compression" help="default: snappy">
      <a-radio-group v-model="data.compression" >
        <a-radio :value="0">none</a-radio>
        <a-radio :value="1">snappy</a-radio>
        <a-radio :value="2">zstd</a-radio>
      </a-radio-group>
    </a-form-item>

    <a-form-item label="syncWrites" help="default: close">
      <a-switch v-model="data.syncWrites" type="round"></a-switch>
    </a-form-item>

    <a-form-item label="gcInterval" help="default: 10 minutes">
      <Duration :data="data.gcInterval" />
      <a-tag class="ml-2" color="orangered">{{$t('needRestartApp')}}</a-tag>
    </a-form-item>

    <a-form-item label="gcDiscardRatio" help="default: 0.9">
      <a-input-number class="w-64" v-model="data.gcDiscardRatio" :step="0.1" :precision="1" :min="0" :max="1" />
      <a-tag class="ml-2" color="orangered">{{$t('needRestartApp')}}</a-tag>
    </a-form-item>

</template>

<script setup>
  import Duration from "@/components/utils/Duration.vue";

  defineProps({data:Object})

  const loadingModeOptions = [
    { label: "FileIO", value: 0 },
    { label: "LoadToRAM", value: 1 },
    { label: "MemoryMap", value: 2 },
  ];

</script>
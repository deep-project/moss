<template>

     <a-form-item label="path">
       <a-input class="w-64" v-model="data.path" />
     </a-form-item>

     <a-form-item label="MemTableSize" help="内存表文件大小 default: 128">
       <a-input-number class="w-64" v-model="data.memTableSize" :min="0">
         <template #suffix>M</template>
       </a-input-number>
     </a-form-item>

     <a-form-item label="BaseTableSize" help="数据表大小 default: 4">
       <a-input-number class="w-64" v-model="data.baseTableSize" :min="0">
         <template #suffix>M</template>
       </a-input-number>
     </a-form-item>

     <a-form-item label="ValueLogFileSize" help="日志文件大小 default: 512">
       <a-input-number class="w-64" v-model="data.valueLogFileSize" :min="0">
         <template #suffix>M</template>
       </a-input-number>
     </a-form-item>

     <a-form-item label="NumMemtables" help="内存表数量 default: 5">
       <a-input-number class="w-64" v-model="data.numMemtables" :min="1"></a-input-number>
     </a-form-item>

     <a-form-item label="NumLevelZeroTables" help="零级表数量 default: 5">
       <a-input-number class="w-64" v-model="data.numLevelZeroTables" :min="1"></a-input-number>
     </a-form-item>

     <a-form-item label="NumLevelZeroTablesStall" help="零级表停滞数量 default: 10">
       <a-input-number class="w-64" v-model="data.numLevelZeroTablesStall" :min="1"></a-input-number>
     </a-form-item>

     <a-form-item label="NumCompactors" help="压缩工数量 default: 4">
       <a-input-number class="w-64" v-model="data.numCompactors" :min="2"></a-input-number>
     </a-form-item>


     <a-form-item label="GcInterval" help="垃圾回收间隔">
       <Duration :data="data.gcInterval" />
       <a-tag class="ml-2" color="orangered">need to restart the app</a-tag>
     </a-form-item>

    <a-form-item label="GcDiscardRatio" help="垃圾回收比例">
       <a-input-number class="w-64" v-model="data.gcDiscardRatio" :step="0.1" :precision="1" :min="0" :max="1" />
       <a-tag class="ml-2" color="orangered">need to restart the app</a-tag>
     </a-form-item>
</template>

<script setup>
  import Duration from "@/components/utils/Duration.vue";

  defineProps({data:Object})

</script>
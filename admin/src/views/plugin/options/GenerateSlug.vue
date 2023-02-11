<template>
  <a-form-item :label="$t('article')">
    <a-select v-model="data.article.style" :options="options" :style="{width:'180px'}"></a-select>
  </a-form-item>

  <a-form-item :label="$t('category')">
    <a-select v-model="data.category.style" :options="options" :style="{width:'180px'}"></a-select>
  </a-form-item>

  <a-form-item :label="$t('tag')">
    <a-select v-model="data.tag.style" :options="options" :style="{width:'180px'}"></a-select>
  </a-form-item>

  <a-table :columns="columns" :data="example" :pagination="false" size="small" />

</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")
 const options = ['snowflake','hashSnowflake', 'idgenerator','hashIdgenerator', 'xid', 'uuid']

 const columns = [{title: 'Name',dataIndex: 'name',width:160}, {title: 'Length',dataIndex: 'length'}, {title: 'Example',dataIndex: 'example'}]
 const example = [
          {name:'snowflake',example:'446936490811851011',length:18},
          {name:'hashSnowflake',example:'VwXWDylVL88P',length:12},
          {name:'idgenerator',example:'384165655576645',length:15},
          {name:'hashIdgenerator',example:'drnKy1WxkP',length:10},
          {name:'xid',example:'cfi9d5bu93k0im7ck2hg',length:20},
          {name:'uuid',example:'1f8cfe73-887f-4823-ba54-bc8a640803f6',length:36},
 ]
</script>

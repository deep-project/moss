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

  <a-table :columns="columns" :data="example" :pagination="false" size="small">

    <template #collision="{ record,rowIndex,column }">
      <a-rate :default-value="record.collision" readonly allow-half class="collision" />
    </template>
  </a-table>

</template>


<script setup>
 import {inject} from "vue";
 const data = inject("options")
 const options = ['snowflake','hashSnowflake', 'idgenerator','hashIdgenerator', 'xid', 'uuid','nanoid',{value:'nanoid8',label:'nanoid-8'}]

 const columns = [
      {title: 'Name', dataIndex: 'name',width:145},
      {title: 'Length', dataIndex: 'length',width:60, sortable: { sortDirections: ['ascend', 'descend'] }},
      {title: 'Example', dataIndex: 'example',width:200},
      {title: 'Anti-collision', dataIndex: 'collision',width:100,slotName:'collision',},

 ]

 const example = [
          {name:'snowflake',example:'446936490811851011',length:18,collision:4},
          {name:'hashSnowflake',example:'VwXWDylVL88P',length:12,collision:4},
          {name:'idgenerator',example:'384165655576645',length:15,collision:3.5},
          {name:'hashIdgenerator',example:'drnKy1WxkP',length:10,collision:3.5},
          {name:'xid',example:'cfi9d5bu93k0im7ck2hg',length:20,collision:3.5},
          {name:'uuid',example:'1f8cfe73-887f-4823-ba54-bc8a640803f6',length:36,collision:5},
          {name:'nanoid',example:'ENqY9KshxBb_PYD5oHY9q',length:21,collision:5},
          {name:'nanoid-8',example:'j2KPn-xk',length:8,collision:2},
 ]
</script>


<style>
.collision {
  font-size: 12px;
}
.collision>.arco-rate-character:not(:last-child){
  margin-right:4px;
}
</style>
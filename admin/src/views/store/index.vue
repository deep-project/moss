<template>
  <Table modelName="store" :columns="columns" order="id desc"  postWidth="600px"  formStyle="padding-right: 10px" formLayout="horizontal" :postComponent="postComponent" />
</template>

<script setup>
import {shallowRef} from 'vue'
import Table from '@/components/dataTable/index.vue'
import { searchFilter } from '@/components/dataTable/index.js'
import Post from './Post.vue'
import {t} from '@/locale'

const postComponent = shallowRef(Post);
const columns = [
  {
    title: t('id'),
    dataIndex: 'id',
    width: 100,
    ellipsis:true,
    filterable: searchFilter,
    sortable: { sortDirections: ['ascend', 'descend'] }
  },
  {
    title: t('title'),
    dataIndex: 'title',
    filterable: searchFilter,
    width: 300,
  },
  {
    title:  t('slug'),
    dataIndex: 'slug',
    filterable: searchFilter,
    width: 200,
    ellipsis:true,
    tooltip:true,
  },
  {
    title: t('category') + ' ID',
    dataIndex: 'category_id',
    width: 100,
    ellipsis:true,
    filterable: searchFilter,
  },
  {
    title: t('category'),
    dataIndex: 'category_name',
    width: 110,
    ellipsis:true,
  },
  {
    title: t('tags'),
    dataIndex: 'tags',
    width: 240,
    ellipsis:true,
    slotName:'tag',
  },
  {
    title: t('createTime'),
    dataIndex: 'store_create_time',
    slotName:'time',
    width: 160,
    align:'right',
  },
  {
    title: t('publish'),
    slotName:'storePost',
    width: 120,
    align:'center',
  },
];
</script>
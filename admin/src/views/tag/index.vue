<template>
  <Table modelName="tag" :columns="columns" order="id desc" postWidth="560px" formStyle="padding-right: 10px" formLayout="horizontal" :postComponent="postComponent" />
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
    title: t('name'),
    dataIndex: 'name',
    filterable: searchFilter,
    width: 300,
    slotName:'title',
  },
  {
    title:  t('slug'),
    dataIndex: 'slug',
    filterable: searchFilter,
    width: 140,
    ellipsis:true,
    tooltip:true,
  },
  {
    title: t('createTime'),
    dataIndex: 'create_time',
    slotName:'time',
    width: 140,
    align:'right',
  },
];
</script>
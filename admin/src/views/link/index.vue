<template>
  <Table modelName="link" :columns="columns" order="id desc" postWidth="560px" formStyle="padding-right: 10px" formLayout="horizontal" :postComponent="postComponent" />
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
    width: 200,
  },
  {
    title:  t('url'),
    dataIndex: 'url',
    filterable: searchFilter,
    slotName:'url',
    width: 300,
    ellipsis:true,
    tooltip:true,
  },
  {
    title:  t('note'),
    dataIndex: 'note',
    width: 250,
    ellipsis:true,
    tooltip:true,
  },
  {
    title: t('expireTime'),
    dataIndex: 'expire_time',
    slotName:'expire_time',
    width: 160,
    align:'right',
  },
  {
    title: t('createTime'),
    dataIndex: 'create_time',
    slotName:'time',
    width: 160,
    align:'right',
  },
  {
    title: t('detect'),
    dataIndex: 'detect',
    slotName:'detect',
    align:'center',
    width: 80,
  },
  {
    title: t('status'),
    dataIndex: 'status',
    slotName:'linkStatus',
    align:'center',
    width: 80,
  },
];
</script>
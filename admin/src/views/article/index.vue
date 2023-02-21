<template>
    <Table modelName="article" :columns="columns" order="id desc" postWidth="98%" postHeight="96%" formLayout="vertical" :postComponent="postComponent" />
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
      width:100,
      ellipsis:true,
      filterable: searchFilter,
      sortable: { sortDirections: ['ascend', 'descend'] }
    },
    {
      title: t('title'),
      dataIndex: 'title',
      filterable: searchFilter,
      width: 300,
      slotName:'title',
      ellipsis:true,
      tooltip:true,
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
      title: t('category'),
      dataIndex: 'category_id',
      width: 100,
      ellipsis:true,
      filterable: searchFilter,
      align:'right',
    },
    {
      title: t('views'),
      dataIndex: 'views',
      width: 100,
      ellipsis:true,
      sortable: { sortDirections: ['ascend', 'descend'] },
      align:'right',
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
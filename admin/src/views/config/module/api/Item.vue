<template>
  <div class="pt-3">
    <a-card v-for="row in data" :title="row.title" class="mb-5">
      <div class="w-full mb-4">
        <a-tag size="large" color="arcoblue">{{row.method}}</a-tag>
        <a-tag size="large" class="mx-2">{{apiURL + row.url}}</a-tag>
        <icon-copy @click="useCopy(apiURL + row.url)" class="cursor-pointer hover:text-blue-500" />
      </div>
      <a-table size="small" v-if="row.payload" :columns="columns" :data="row.payload" :pagination="false">
          <template #required="{ record,rowIndex }">
           <icon-check-circle-fill v-if="record.required" style="color:rgb(var(--arcoblue-6))" />
           <span v-else> - </span>
          </template>
      </a-table>

    </a-card>
  </div>
</template>

<script setup>

  import {useCopy} from "@/hooks/utils";
  import {useStore} from "@/store";
  import {computed, h} from "vue";
  import {t} from "@/locale/index.js";
  import {useAppendSiteURL} from "@/hooks/app/index.js";

  defineProps({data:Object})
  const store = useStore()
  let apiURL = useAppendSiteURL(store, store.config["router"].admin_path + "/api")

  const columns = computed(()=> [
    { title: t('fields'), dataIndex: 'field', width:150 },
    { title: t('type'), dataIndex: 'type', width:150 },
    { title: t('required'), dataIndex: 'required', width:150,slotName:'required'},
    { title: t('description'), dataIndex: 'description',width:300},
    { title: t('example'), dataIndex: 'example',width:300},
  ])

</script>
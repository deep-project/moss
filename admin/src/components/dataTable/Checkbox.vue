<template>
  <template v-if="selectedKeys.length > 0">
    <a-dropdown :popup-max-height="false" trigger="hover">
      <a-button size="small">
        <template #icon><icon-down /></template>
        {{ $t('action') }}
      </a-button>
      <template #content>
        <a-doption @click="batchDelete(modelName, selectedKeys)" style="color: rgb(var(--red-6))">
          <template #icon><icon-delete /></template>
          {{$t('delete')}}
        </a-doption>
      </template>
    </a-dropdown>
    <a-tag class="ml-4 overflow-hidden" style="border-radius:18px;background-color: var(--color-fill-2)"
           @close="selectedKeys = []" closable>{{$t('checked')}} {{ selectedKeys.length }}</a-tag>
  </template>
</template>


<script setup>
  import {inject} from 'vue'

  const selectedKeys = inject('selectedKeys')
  const batchDelete = inject('batchDelete')
  const modelName = inject('modelName')

</script>
<template>
  <a-button-group size="mini">
    <a-button @click="onEdit(record,rowIndex)">{{$t('edit')}}</a-button>
    <a-dropdown position="br" size="small">
      <a-button><template #icon><icon-down /></template></a-button>
      <template #content>
        <a-doption @click="runDelete(modelName, record.id)" style="color: rgb(var(--red-6))">
          <template #icon><icon-delete /></template>
          {{$t('delete')}}
        </a-doption>
        <!-- <a-divider :margin="4" />-->
      </template>
    </a-dropdown>
  </a-button-group>
</template>

<script setup>
  defineProps({ record:Object,modelName:String, rowIndex:Number, runDelete:Function, onEdit:Function })
</script>
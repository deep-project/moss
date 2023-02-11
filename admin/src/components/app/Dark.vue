<template>
  <a-switch type="line" v-model="store.dark">
    <template #checked-icon>
      <icon-moon-fill class="icon" />
    </template>
    <template #unchecked-icon>
      <icon-sun-fill class="icon" :style="{color:store.color}" />
    </template>
  </a-switch>
</template>

<script setup>
  import {useStore} from "@/store";

  const store = useStore()
</script>
<style scoped>
.icon{
  transition: all .2s;
  transform: rotateZ(-90deg);
}
.icon:hover{
  transform: rotateZ(0deg);
}
</style>
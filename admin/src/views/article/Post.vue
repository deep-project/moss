<template>
  <a-layout class="h-full">
    <a-layout>
      <a-layout-content class="h-full" style="padding:0 12px 0 8px">
        <PostLeft />
      </a-layout-content>
    </a-layout>
    <PostRight />
  </a-layout>
</template>


<script setup>

  import PostRight from "@/views/article/PostRight.vue";
  import PostLeft from "@/views/article/PostLeft.vue";

</script>
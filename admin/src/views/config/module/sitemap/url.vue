<template>
  <a-space direction="vertical">
    <a-divider type="dashed" :margin="10" orientation="left">xml</a-divider>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/article.xml</a-tag>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/category.xml</a-tag>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/tag.xml</a-tag>
    <a-divider type="dashed" :margin="10" orientation="left">txt</a-divider>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/article.txt</a-tag>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/category.txt</a-tag>
    <a-tag size="small" class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{sitemapURL}}/tag.txt</a-tag>
  </a-space>
</template>



<script setup>

  import {useStore} from "@/store";
  import {useOpenLink} from "@/hooks/utils.js";
  import {computed} from "vue";

  const store = useStore()
  const {sitemap_path} = defineProps({sitemap_path:String})
  const confSitemapPath = store.config['router'].sitemap_path
  let siteURL = store.config['site'].url
  if(!siteURL) siteURL = window.location.origin

  const sitemapURL = computed(()=>{
    let dir = sitemap_path ? sitemap_path:confSitemapPath
    if(!dir) dir = "/sitemap"
    if(dir.indexOf("/") !== 0) dir = "/" + dir

    return siteURL + dir
  })

</script>


<style scoped>
.cursor-pointer{
  cursor: pointer;
}
</style>
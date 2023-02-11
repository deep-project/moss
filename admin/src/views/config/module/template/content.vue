<template>
  <a-form-item label="Head" field="head">
    <a-textarea v-model="data.head" :auto-size="{minRows:5, maxRows:20}" />
  </a-form-item>
  <a-form-item label="Footer" field="footer">
    <a-textarea v-model="data.footer" :auto-size="{minRows:5, maxRows:20}" />
  </a-form-item>
  <a-form-item label="robots.txt" field="robots_txt">
    <a-textarea v-model="data.robots_txt" :auto-size="{minRows:5, maxRows:20}" />
    <template #extra><span class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{ siteURL }}/robots.txt</span></template>
  </a-form-item>
  <a-form-item label="ads.txt" field="ads_txt">
    <a-textarea v-model="data.ads_txt" :auto-size="{minRows:5, maxRows:20}" />
    <template #extra><span class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{ siteURL }}/ads.txt</span></template>
  </a-form-item>
  <a-form-item label="favicon.ico" field="favicon_ico">
    <a-textarea v-model="data.favicon_ico" :auto-size="{minRows:4, maxRows:4}" />
    <template #extra>
      <div class="flex items-center gap-3">
        <a-upload :custom-request="faviconIcoBase64" size="small" :show-file-list="false" accept="image/*">
          <template #upload-button><a-button type="outline" size="mini">{{ $t('upload') }}</a-button></template>
        </a-upload>
        <span class="cursor-pointer hover:text-blue-400" @click="useOpenLink">{{ siteURL }}/favicon.ico</span>
      </div>
    </template>
  </a-form-item>
</template>

<script setup>
  import {inject} from 'vue'
  import {useStore} from "@/store";
  import {useOpenLink} from "@/hooks/utils.js";

  const data = inject('data')
  const store = useStore()
  let siteURL = store.config['site'].url
  if(!siteURL) siteURL = window.location.origin


  const faviconIcoBase64 = (option) => {
    const {fileItem} = option
    const reader = new FileReader()
    reader.readAsDataURL(fileItem.file)
    reader.onload = () => {
      data.value.favicon_ico = reader.result
    }
  }

</script>
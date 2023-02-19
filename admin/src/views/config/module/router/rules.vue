<template>
  <a-form-item :label="$t('article')">
    <a-popover position="rt" title="examples:">
      <a-input class="w-64" v-model="data.article_rule" placeholder="/article/{slug}" />
      <template #content>
        <a-space direction="vertical">
          <a-tag size="small" class="cursor-pointer" @click="useCopy">/article/{slug}</a-tag>
          <a-tag size="small" class="cursor-pointer" @click="useCopy">/archives/{slug}.html</a-tag>
          <a-tag size="small" class="cursor-pointer" @click="useCopy">/post-{slug}</a-tag>
        </a-space>
      </template>
    </a-popover>
  </a-form-item>

  <a-form-item :label="$t('category')">
    <a-space class="w-full" direction="vertical">
      <a-popover position="rt" title="examples:">
        <a-input class="w-64" v-model="data.category_rule" placeholder="/category/{slug}" />
        <template #content>
          <a-space direction="vertical">
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/category/{slug}</a-tag>
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/{slug}.html</a-tag>
          </a-space>
        </template>
      </a-popover>
      <a-popover position="rt" title="examples:">
        <a-input class="w-64" v-model="data.category_page_rule" placeholder="/category/{slug}/{page}" />
        <template #content>
          <a-space direction="vertical">
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/category/{slug}/{page}</a-tag>
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/{slug}/{page}.html</a-tag>
          </a-space>
        </template>
      </a-popover>
    </a-space>
  </a-form-item>

  <a-form-item :label="$t('tag')">
    <a-space class="w-full" direction="vertical">
      <a-popover position="rt" title="examples:">
        <a-input class="w-64" v-model="data.tag_rule" placeholder="/tag/{slug}" />
        <template #content>
          <a-space direction="vertical">
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/tag/{slug}</a-tag>
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/{slug}.html</a-tag>
          </a-space>
        </template>
      </a-popover>
      <a-popover position="rt" title="examples:">
        <a-input class="w-64" v-model="data.tag_page_rule" placeholder="/tag/{slug}/{page}" />
        <template #content>
          <a-space direction="vertical">
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/tag/{slug}/{page}</a-tag>
            <a-tag size="small" class="cursor-pointer" @click="useCopy">/tag-{slug}/page-{page}.html</a-tag>
          </a-space>
        </template>
      </a-popover>
    </a-space>
  </a-form-item>

  <div style="max-width: 527px">
    <a-alert class="mb-2">
      <a-space>
        <span>{{ $t('delimiter') }}:</span>
        <a-tag>/</a-tag>
        <a-tag>-</a-tag>
        <a-tag>.</a-tag>
        <a-tag>:</a-tag>
      </a-space>
    </a-alert>

    <a-alert style="max-width:700px" type="warning">
      如果别名中可能包含某个分隔符，则路由中避免使用此分隔符，否则无法匹配路由。建议使用<a-tag class="ml-2">/</a-tag>
    </a-alert>
  </div>

  <a-divider type="double"></a-divider>

  <a-form-item :label="$t('adminPath')">
    <a-input class="w-64" v-model="data.admin_path" placeholder="/admin" />
  </a-form-item>

  <a-form-item :label="$t('sitemap')">
    <a-popover position="rt">
      <a-input class="w-64" v-model="data.sitemap_path" placeholder="/sitemap" />
      <template #content>
        <SitemapURL :sitemap_path="data.sitemap_path" />
      </template>
    </a-popover>
  </a-form-item>

</template>


<script setup>
  import {useCopy} from "@/hooks/utils.js";
  import {inject} from "vue";
  import SitemapURL from '../sitemap/url.vue'

  const data = inject('data')
</script>
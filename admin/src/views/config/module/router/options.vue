<template>

    <a-form-item label="Etag">
      <a-switch type="round" v-model="data.etag" />
    </a-form-item>

  <a-form-item label="Minify code">
    <a-switch type="round" v-model="data.minify_code" />
    <a-tooltip>
      <icon-question-circle class="ml-2 opacity-60" />
      <template #content>Minify HTML/CSS/JS code. Remove blank lines.</template>
    </a-tooltip>
  </a-form-item>

  <a-form-item :label="$t('compress')">
      <a-select class="w-64" v-model="data.compress_level" :options="optionsCompressLevel" />
        <a-tooltip>
          <icon-question-circle class="ml-2 opacity-60" />
          <template #content>Gzip,Brotli compress</template>
        </a-tooltip>
    </a-form-item>

    <a-form-item label="Proxy header">
      <a-select class="w-64" v-model="data.proxy_header" multiple allow-clear allow-create :options="optionsProxyHeader" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>If the request from some sort of proxy, like a load balancer. set to get the correct ip.</template>
      </a-tooltip>
    </a-form-item>

    <a-divider />

    <a-form-item label="Body limit">
      <a-input-number class="w-64" v-model="data.options.body_limit" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            Sets the maximum allowed size for a request body, if the size exceeds the configured limit, it sends 413 - Request Entity Too Large response.
            <br>-1 will decline any body size
            <br>0 Default: 4 * 1024 * 1024
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Concurrency">
      <a-input-number class="w-64" v-model="data.options.concurrency" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>Maximum number of concurrent connections.<br>0 Default: 256 * 1024</template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Read buffer size">
      <a-input-number class="w-64" v-model="data.options.read_buffer_size" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            Per-connection buffer size for requests' reading. This also limits the maximum header size.
            <br>Increase this buffer if your clients send multi-KB RequestURIs and/or multi-KB headers (for example, BIG cookies).
            <br>0 Default: 4096
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Write buffer size">
      <a-input-number class="w-64" v-model="data.options.write_buffer_size" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>Per-connection buffer size for responses' writing.<br>0 Default: 4096</template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Server header">
      <a-input v-model="data.options.server_header" class="w-64" placeholder="" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>Enables the "Server: value" HTTP header.<br>Default: ""</template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Network">
      <a-select class="w-64" v-model="data.options.Network" :options="optionsNetwork" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
          Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only)
          <br>WARNING: When prefork is set to true, only "tcp4" and "tcp6" can be chose.
          <br>Default: tcp4
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Strict routing">
      <a-switch type="round" v-model="data.options.strict_routing" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, the router treats "/foo" and "/foo/" as different.<br>
            By default this is disabled and both "/foo" and "/foo/" will execute the same handler.<br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Case sensitive">
      <a-switch type="round" v-model="data.options.case_sensitive" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, enables case sensitive routing.<br>
            E.g. "/FoO" and "/foo" are treated as different routes.<br>
            By default this is disabled and both "/FoO" and "/foo" will execute the same handler.<br>
            Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <!--    <n-form-item label="Prefork">-->
    <!--      <n-switch :round="false" v-model:value="data.options.prefork" />-->
    <!--      <n-tooltip trigger="hover">-->
    <!--        <template #trigger><n-icon class="ml-2" :depth="4" size="18"><HelpCircleOutline /></n-icon></template>-->
    <!--        When set to true, this will spawn multiple Go processes listening on the same port.<br>Default: false-->
    <!--      </n-tooltip>-->
    <!--    </n-form-item>-->

    <!--    <n-form-item label="Get only">-->
    <!--      <n-switch :round="false" v-model:value="data.options.get_only" />-->
    <!--      <n-tooltip trigger="hover" :style="{ maxWidth: '600px' }">-->
    <!--        <template #trigger><n-icon class="ml-2" :depth="4" size="18"><HelpCircleOutline /></n-icon></template>-->
    <!--        GETOnly rejects all non-GET requests if set to true.-->
    <!--        <br>This option is useful as anti-DoS protection for servers accepting only GET requests.-->
    <!--        <br>The request size is limited by ReadBufferSize if GETOnly is set.-->
    <!--        <br>Default: false-->
    <!--      </n-tooltip>-->
    <!--    </n-form-item>-->

    <a-form-item label="Disable keepalive">
      <a-switch type="round" v-model="data.options.disable_keepalive" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            Disable keep-alive connections, the server will close incoming connections after sending the first response to the client
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Disable default date">
      <a-switch type="round" v-model="data.options.disable_default_date" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, causes the default date header to be excluded from the response.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Disable default content type">
      <a-switch type="round" v-model="data.options.disable_default_content_type" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, causes the default Content-Type header to be excluded from the response.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Disable header normalizing">
      <a-switch type="round" v-model="data.options.disable_header_normalizing" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, disables header normalization.
            <br>By default all header names are normalized: conteNT-tYPE -> Content-Type.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Disable startup message">
      <a-switch type="round" v-model="data.options.disable_startup_message" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            When set to true, it will not print out the «Fiber» ASCII art and listening address.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Stream request body">
      <a-switch type="round" v-model="data.options.StreamRequestBody" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            StreamRequestBody enables request body streaming, and calls the handler sooner when given body is larger then the current limit.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

    <a-form-item label="Reduce memory usage">
      <a-switch type="round" v-model="data.options.reduce_memory_usage" />
      <a-tooltip>
        <icon-question-circle class="ml-2 opacity-60" />
        <template #content>
            Aggressively reduces memory usage at the cost of higher CPU usage if set to true.
            <br>Try enabling this option only if the server consumes too much memory serving mostly idle keep-alive connections.
            This may reduce memory usage by more than 50%.
            <br>Default: false
        </template>
      </a-tooltip>
    </a-form-item>

  <a-alert type="warning">{{$t('needRestartApp')}}</a-alert>

</template>


<script setup>
  import {useStore} from "@/store";
  import {t} from "@/locale";
  import {inject,computed} from "vue";

  const data = inject('data')
  const store = useStore()

  const optionsNetwork = computed(()=>{
    return [
      { label: "tcp", value: 'tcp' },
      { label: "tcp4 (IPv4-only)", value: 'tcp4' } ,
      { label: "tcp6 (IPv6-only) "+ t('warning')+'!!', value: 'tcp6' } ,
    ]
  }) ;

  const optionsProxyHeader = [
    { label: 'X-Forwarded-For', value: 'X-Forwarded-For' },
    { label: 'CF-Connecting-IP', value: 'CF-Connecting-IP' },
    { label: 'Ali-Cdn-Real-Ip', value: 'Ali-Cdn-Real-Ip' },
  ]
  const optionsCompressLevel = computed(()=> [
    { label: t('disable'), value: -1 },
    { label: t('default'), value: 0 },
    { label: t('bestSpeed'), value: 1},
    { label: t('bestCompression'), value: 2},
  ])

</script>
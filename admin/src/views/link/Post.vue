<template>
  <a-form-item field="name" :label="$t('name')" :rules="[{ required:true, message: $t('message.required',[$t('name')]) }]">
    <a-input v-model="record.name" :max-length="150" allow-clear show-word-limit />
  </a-form-item>

  <a-form-item field="url" :label="$t('url')" :rules="[{ required:true, message: $t('message.required',[$t('url')]) }]">
    <a-input v-model="record.url" :max-length="250" allow-clear show-word-limit />
  </a-form-item>

  <a-form-item field="logo" label="logo">
    <UploadImgInput class="w-full" v-model="record.logo" />
  </a-form-item>

  <a-form-item field="note" :label="$t('note')">
    <a-textarea class="input" v-model="record.note" :max-length="250" :auto-size="{minRows:3,maxRows:5}" show-word-limit />
  </a-form-item>

  <a-form-item field="expire_time" :label="$t('expireTime')">
    <a-date-picker class="w-full" v-model="expireTime" value-format="timestamp" show-time @change="(val)=>record.expire_time =parseInt(val / 1000)"></a-date-picker>
    <a-button @click="incrExpireTime"><template #icon><a-link :hoverable="false"><icon-plus-circle /></a-link></template></a-button>
  </a-form-item>

  <a-form-item v-if="record.id > 0" field="create_time" :label="$t('createTime')">
    <a-date-picker class="w-full" v-model="createTime" value-format="timestamp" show-time @change="(val)=>record.create_time =parseInt(val / 1000)" />
  </a-form-item>

  <a-form-item field="detect" :label="$t('detect')" help="detect it contains my link by scheduled task">
    <a-space>
      <a-switch type="round" v-model="record.detect" />
      <template v-if="record.detect">
        <a-input-number size="small" v-model="record.detect_delay" :min="0" style="width:180px">
          <template #prepend><span style="font-size:12px">{{$t('delay')}}</span></template>
        </a-input-number>
        <span class="text-gray-500" style="font-size: 12px">{{$t('minutes')}}</span>
      </template>


    </a-space>
  </a-form-item>


  <a-form-item field="status" :label="$t('status')">
    <a-switch type="round" v-model="record.status" />
  </a-form-item>

</template>


<script setup>
  import {computed, inject} from "vue";
  import UploadImgInput from '@/components/utils/UploadImgInput.vue'


  const record = inject('record')
  const expireTime = computed(()=>record.value.expire_time*1000)
  const createTime = computed(()=>record.value.create_time*1000)
  if(!record.value.id){
    record.value.detect = true
    record.value.status = true
    record.value.detect_delay = 60
  }

  let incrExpire = false
  function incrExpireTime(){
    if(!record.value.id && incrExpire === false){
      record.value.detect = false // 创建时第一次添加过期时间，主动把检查关闭
    }
    if(!expireTime.value) record.value.expire_time = new Date().getTime() / 1000
    record.value.expire_time += 60 * 60 * 24 * 31
    incrExpire = true
  }

</script>
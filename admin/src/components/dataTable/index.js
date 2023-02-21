import {h} from "vue";
import {IconSearch} from "@arco-design/web-vue/es/icon/index.js";
import {Message} from "@arco-design/web-vue";
import {t} from '@/locale'

export const searchFilter = {
    slotName: 'column-search',
    icon: () => h(IconSearch)
}



export function paginationOption(loadingCount, total, params, store){
    return {
        total:loadingCount.value ? 9999 : total.value,
            defaultPageSize:params.limit,
            showTotal:!loadingCount.value,
            showJumper:!store.isMobile && !loadingCount.value,
            showPageSize:!store.isMobile && !loadingCount.value,
            simple:store.isMobile || loadingCount.value,
            pageSizeOptions:[ 10, 20, 30, 40, 50, 100, 500, 1000, 2000, 5000 ],
    }
}

export function listOption(modelName,params){
    return {
        defaultParams:[modelName, params],
        //loadingKeep: 600
    }
}

export function deleteOption(refreshList){
    return {
        manual:true,
        onSuccess:(resp)=>{
            if(resp.success) refreshList()
        }
    }
}

export function batchDeleteOption(selectedKeys,refresh){
    return {
        manual:true,
        onSuccess:(resp)=>{
            if(!resp.success) return
            selectedKeys.value = []
            refresh()
        }
    }
}

export function getOption(postRecord,onSuccessCallback){
    return {
        manual:true,
        onSuccess:(data)=>{
            postRecord.value = data
            if(onSuccessCallback.value) onSuccessCallback.value(data)
        }
    }
}

export function createOption(visiblePost, refresh, beforeCallback){
    return  {
        manual:true,
        onBefore:()=>{
            if(beforeCallback.value) beforeCallback.value()
        },
        onSuccess:(resp)=>{
            if(!resp.success) return postFail(resp)
            visiblePost.value = false
            refresh()
        }
    }
}

export function updateOption(visiblePost, data, rowIndex, postRecord,successCallback){
    return {
        manual:true,
        onSuccess:(resp)=>{
            if(!resp.success) return postFail(resp)
            visiblePost.value = false
            for(let key in data.value[rowIndex.value]){
                data.value[rowIndex.value][key] = postRecord.value[key]
            }
            if(successCallback.value) successCallback.value()
        }
    }
}

export function storePostOption(refreshList){
    return {
        manual:true,
        onSuccess:(resp)=>{
            if(resp.success) {
                Message.success(t('message.success',[t('publish')]))
                refreshList()
            }
        }
    }
}

function postFail(resp){
    if(resp.message.toLowerCase().indexOf("unique") > -1 && resp.message.toLowerCase().indexOf("slug") > -1){
        Message.error(t('message.exists',[t('slug')]))
    }
}
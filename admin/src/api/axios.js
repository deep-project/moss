
import axios from "axios"
import {useStore} from "@/store";
import { Message } from '@arco-design/web-vue';

const _axios = axios.create({
    baseURL: process.env.NODE_ENV === "production" ? "/{{__DIR__}}/api" : "/admin/api",
});

// 请求拦截
_axios.interceptors.request.use(
    function(config) {
        config.headers['token'] = useStore().token
        // 给get请求加时间戳，防止被CDN缓存
        if(config.method === "get"){
            let time = (new Date()).valueOf()
            if(config.url.indexOf("?") === -1) config.url = config.url + "?t=" + time
            else config.url = config.url + "&t=" + time
        }
        return config;
    },
    function(error) {
        if(error.message) Message.error({content:error.message,resetOnHover:true})
        return Promise.reject(error);
    }
);

var count_401 = 0

// 响应拦截
_axios.interceptors.response.use(
    function(response) {
        let success = response.data.success
        let msg = response.data.message
        if(msg){
            success ? Message.success(msg) : Message.error({content:msg,resetOnHover:true})
        }
        return response;
    },
    function(error) {
        if(error.response.status===401) {
            count_401++
            const store = useStore()
            store.token = ""
        }
        if(count_401 >= 2) setTimeout(() =>{ count_401 = 0},4000) // 4秒后清零
        if(error.message && count_401 <= 2) Message.error({content:error.message,resetOnHover:true})
        return Promise.reject(error);
    }
);

export default _axios;
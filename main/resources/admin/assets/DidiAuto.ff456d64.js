import{b as y}from"./index.97dfccda.js";import{i as b,j as n,k as _,p as x,v as l,m as r,$ as t,y as s,n as d,z as u,l as U,s as k,F as v}from"./@vue.0987707a.js";import"./vue-router.0ed66d6f.js";import"./vue-i18n.e3137642.js";import"./@intlify.bed9fa1a.js";import"./source-map.205bdfab.js";import"./vue.5c5bb0aa.js";import"./@arco-design.38d3ffc1.js";import"./resize-observer-polyfill.8deb1e21.js";import"./compute-scroll-into-view.17358474.js";import"./b-tween.87ffe365.js";import"./dayjs.396bdce9.js";import"./b-validate.ee581f7d.js";import"./number-precision.6dad9ff9.js";import"./scroll-into-view-if-needed.61c672a4.js";import"./@vueuse.d5398ce4.js";import"./pinia.2e07300c.js";import"./vue-demi.b3a9cad9.js";import"./naive-ui.0057ea16.js";import"./date-fns.e2bf381f.js";import"./seemly.d0f7d7a4.js";import"./evtd.9eee5233.js";import"./@css-render.6ced7bf3.js";import"./css-render.20ab466e.js";import"./@emotion.6322e2ae.js";import"./vooks.3f61458b.js";import"./vueuc.5f5811a3.js";import"./vdirs.ab69c576.js";import"./@juggle.32c34d6c.js";import"./lodash-es.33d1f95f.js";import"./date-fns-tz.6b0f78d1.js";import"./axios.b9f958b0.js";import"./vue-request.6886b8d7.js";/* empty css                    */const g={class:"text-sm text-gray-400 ml-3"},$={class:"text-sm text-gray-400 ml-3"},B={class:"text-sm text-gray-400 ml-3"},D={__name:"DidiAuto",setup(N){const e=b("options");return(p,o)=>{const i=n("a-input-number"),m=n("a-form-item"),c=n("a-textarea"),V=n("a-divider"),f=n("a-input");return _(),x(v,null,[l(m,{label:"\u7AD9\u70B9id",required:""},{default:r(()=>[l(i,{modelValue:t(e).site_id,"onUpdate:modelValue":o[0]||(o[0]=a=>t(e).site_id=a),class:"input",min:0},null,8,["modelValue"])]),_:1}),l(m,{label:"cookie",required:""},{default:r(()=>[l(c,{modelValue:t(e).cookie,"onUpdate:modelValue":o[1]||(o[1]=a=>t(e).cookie=a)},null,8,["modelValue"])]),_:1}),l(V),l(m,{label:"\u5EF6\u8FDF\u68C0\u6D4B\u65F6\u95F4"},{default:r(()=>[l(i,{modelValue:t(e).detect_delay,"onUpdate:modelValue":o[2]||(o[2]=a=>t(e).detect_delay=a),class:"input",min:0},null,8,["modelValue"]),s(),d("span",g,u(p.$t("minutes")),1)]),_:1}),l(m,{label:p.$t("retry")},{default:r(()=>[l(i,{modelValue:t(e).retry,"onUpdate:modelValue":o[3]||(o[3]=a=>t(e).retry=a),class:"input",min:0},null,8,["modelValue"])]),_:1},8,["label"]),l(m,{label:"\u91CD\u8BD5\u7B49\u5F85"},{default:r(()=>[l(i,{modelValue:t(e).retry_sleep,"onUpdate:modelValue":o[4]||(o[4]=a=>t(e).retry_sleep=a),class:"input",min:0},null,8,["modelValue"]),s(),d("span",$,u(p.$t("seconds")),1)]),_:1}),l(m,{label:"\u6E05\u9664\u7F13\u5B58URL",help:"\u6E05\u9664\u9996\u9875\u7F13\u5B58\u89E6\u53D1\u7684URL"},{default:r(()=>[l(f,{modelValue:t(e).clear_cache_url,"onUpdate:modelValue":o[5]||(o[5]=a=>t(e).clear_cache_url=a),class:"input"},null,8,["modelValue"])]),_:1}),t(e).clear_cache_url?(_(),U(m,{key:0,label:"\u6E05\u9664\u7F13\u5B58\u540E\u7B49\u5F85"},{default:r(()=>[l(i,{modelValue:t(e).clear_cache_sleep,"onUpdate:modelValue":o[6]||(o[6]=a=>t(e).clear_cache_sleep=a),class:"input",min:0},null,8,["modelValue"]),s(),d("span",B,u(p.$t("seconds")),1)]),_:1})):k("",!0)],64)}}},ne=y(D,[["__scopeId","data-v-f8ebf9c2"]]);export{ne as default};
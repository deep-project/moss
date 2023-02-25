import {Message} from "@arco-design/web-vue";
import {t} from '@/locale'
// 深拷贝
export function useDeepCopy(obj) {
    let _obj = JSON.stringify(obj);
    return JSON.parse(_obj);
}


// 随机字符串
export function useRandString(len, chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678'){
    let res = "";
    let count = chars.length;
    for(let i = 0; i < len ; i ++) {
        res += chars.charAt(Math.floor(Math.random() * count));
    }
    return res;
}

// 无痕打开链接
export function useOpenLink(link, title){
    if(typeof link === 'object') link = link.target.innerText
    if(!link && link.path) link = link.path[0].innerText
    let win = window.open('about:blank');
    if(!win) return;
    if(!title) title = 'loading...'
    win.opener = null;  // 防止代码对上级页面修改
    let iframe = document.createElement('iframe')
    iframe.src= 'javascript:"<script>top.location.replace(`'+link+'`)<//script>"'.replace('//','/')
    iframe.style.display = 'none';
    win.document.write('<title>'+title+'</title>')
    win.document.write('<meta name="referrer" content="never"><meta name="referrer" content="never"><link rel="stylesheet" href="/loading.css">')
    win.document.write('<style>*{padding:0;margin:0;overflow: hidden;text-align: center;line-height: 28px;font-size:14px}</style>')
    win.document.write(`
<div style="width:100vw;height:100vh;display:flex;justify-content: center;align-items: center;">
    <div class="app-loading-container">
    <div class="app-loading">
    <div class="shape1"></div>
    <div class="shape2"></div>
    <div class="shape3"></div>
    <div class="shape4"></div>
    </div>
</div><div style="position: absolute">`+title+'<p>'+link+`</p></div></div>
`)
    setTimeout(() => {
        if(!win) return;
        win.document.body.appendChild(iframe);
    }, 600);
    win.document.close();
}

// copy 复制
export function useCopy(val){
    if(typeof val === 'object') val = val.target.innerText
    let input = document.createElement('input');
    document.body.appendChild(input)
    input.value =  val;
    input.select();
    document.execCommand('copy')
    input.remove();
    Message.success(t('message.success',[t('copy')]));
}


export function useParseBytesSize(bytes){
    if(!bytes || bytes < 1024) return [bytes, "B"] // 不足1K
    if(bytes < 1024*1024) return [bytes / 1024, "K"] // 不足1M
    if(bytes < 1024*1024*1024) return [bytes / 1024 / 1024, "M"] // 不足1G
    if(bytes < 1024*1024*1024*1024) return [bytes / 1024 / 1024 / 1024, "G"] // 不足1G
    if(bytes < 1024*1024*1024*1024*1024) return [bytes / 1024 / 1024 / 1024 / 1024, "T"] // 不足1P
    return [bytes / 1024 / 1024 / 1024 / 1024 / 1024, "P"]
}


export function useIsMobile() {
    if (window.screen.height < 750) {
        return true
    }
    let mobileAgents = ["Android", "iPhone", "SymbianOS", "Windows Phone", "iPod"];
    for (let v = 0; v < mobileAgents.length; v++) {
        if (navigator.userAgent.indexOf(mobileAgents[v]) > 0) {
            return true
        }
    }
    return false;
}
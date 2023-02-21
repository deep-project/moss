
import {useWindowSize} from "@vueuse/core";

export function useDefaultSiderCollapsed(){
    const {width} = useWindowSize()
    return width.value <= 1024;
}

export function useNavigatorDark(){
    return window.matchMedia('(prefers-color-scheme: dark)').matches === true
}


export function useSiteURL(store){
    let u = store.config.site.url
    if (!u) u = window.location.origin
    if (u.endsWith('/')) u = u.slice(0, -1);
    return u
}

export function useAppendSiteURL(store,path){
    if(path.indexOf('/')!==0) path = "/" + path
    return useSiteURL(store) + path
}
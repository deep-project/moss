
import {useWindowSize} from "@vueuse/core";

export function useDefaultSiderCollapsed(){
    const {width} = useWindowSize()
    return width.value <= 1024;
}

export function useNavigatorDark(){
    return window.matchMedia('(prefers-color-scheme: dark)').matches === true
}

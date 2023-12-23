
import {useStorage,useWindowSize} from "@vueuse/core";
import {useNavigatorLocale,la,laid} from "@/locale";
import {defineStore} from "pinia";
import {useDefaultSiderCollapsed,useNavigatorDark} from "@/hooks/app";
import {ref} from 'vue'
import {useIsMobile} from "@/hooks/utils.js";

export const useStore = defineStore('default', {
    state: () => ({
        locale: useStorage("locale", useNavigatorLocale()),
        token: useStorage("token", ""),
        siderCollapsed : useStorage("sider_collapsed", useDefaultSiderCollapsed()),
        siderWidth: useStorage("sider_width", 180), // , localStorage, { listenToStorageChanges: false }
        windowSize:useWindowSize(),
        dark: useStorage("dark", useNavigatorDark()),
        color: useStorage("color", "#4CA1F7"),
        bgColor: useStorage("bg_color", "#e5e4e4"),
        darkBgColor:useStorage("dark_bg_color", "#21252B"),
        mainHeight:ref(0),
        la: la,
        laid:laid,
        config: {},
    }),
    getters: {
        isMobile: (state) => state.windowSize.width < 750 || useIsMobile(),
        isTablet: (state) => state.windowSize.width > 640 && state.windowSize.width <= 1024,
        headerHeight() {
            if (this.isMobile) return 46
            if (this.windowSize.width > 1536) return 60
            return 50
        },
        siderCollapsedWidth: (state) =>{
            if (state.isMobile) return 0
            return 60
        }
    },
    actions:{
        initConfig(resp){
            if(!resp.success) return
            let res = {}
            for(let item of resp.data){
                res[item.id]= item.data
            }
            this.config = res
        }
    }
})
import { createI18n } from 'vue-i18n'

import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';

import en_US from "./lang/en-us";
import zh_CN from "./lang/zh-cn";

import {
    enUS as naiveEnUS,
    zhCN as naiveZhUS,
    dateZhCN,
    dateEnUS,
} from 'naive-ui'

import {useStore} from "@/store/index.js";
import {useNavigatorLanguage} from "@vueuse/core";

export const defaultLocale = "en-us"

export const localeList = {
    'en-us': {name:"English",lang:en_US, arcoLang:enUS, naiveLang: naiveEnUS, naiveDateLang: dateEnUS},
    'zh-cn': {name:"简体中文",lang:zh_CN, arcoLang:zhCN, naiveLang: naiveZhUS, naiveDateLang: dateZhCN},
}

// vue-i18n package
export const i18n = createI18n({
    locale: defaultLocale,
    fallbackLocale: defaultLocale,
    messages:messages()
})

function messages(){
    let arr = {}
    for (let key in localeList){
        arr[key] = localeList[key].lang
    }
    return arr
}


export const t = i18n.global.t
export const tc = i18n.global.tc // 复数转换
export const td = i18n.global.d // 日期转换

// 初始化语言场景
export function useInitLocale(){
    const store = useStore()
    useSetI18nLocale(store.locale)
}

// 设置语言场景
export function useSetLocale(locale){
    const store = useStore()
    store.locale = locale
    useSetI18nLocale(locale)
}

// 设置i18n包语言场景
export function useSetI18nLocale(locale){
    i18n.global.locale = locale
    i18n.global.fallbackLocale = locale
}

// 获取语言场景列表
export function useLocaleList(){
    return localeList
}

// 获取浏览器语言场景
export function useNavigatorLocale(){
    const navLang = useNavigatorLanguage()
    if(navLang.isSupported.value && navLang.language.value) {
        let locale = navLang.language.value.toLowerCase()
        if(useExistsLocale(locale)) return locale
    }
    return defaultLocale
}

export function useExistsLocale(locale){
    return !!localeList[locale];
}

import { createI18n } from 'vue-i18n'

import en_US from "./lang/en-us";
import de_DE from "./lang/de-de";
import es_ES from "./lang/es-es";
import fr_FR from "./lang/fr-fr";
import id_ID from "./lang/id-id";
import it_IT from "./lang/it-it";
import ja_JP from "./lang/ja-jp";
import ko_KR from "./lang/ko-kr";
import pt_BR from "./lang/pt-br";
import th_TH from "./lang/th-th";
import zh_CN from "./lang/zh-cn";
import zh_TW from "./lang/zh-tw";



import enUS from '@arco-design/web-vue/es/locale/lang/en-us';
import deDE from '@arco-design/web-vue/es/locale/lang/de-de';
import esES from '@arco-design/web-vue/es/locale/lang/es-es';
import frFR from '@arco-design/web-vue/es/locale/lang/fr-fr';
import idID from '@arco-design/web-vue/es/locale/lang/id-id';
import itIT from '@arco-design/web-vue/es/locale/lang/it-it';
import jaJP from '@arco-design/web-vue/es/locale/lang/ja-jp';
import koOR from '@arco-design/web-vue/es/locale/lang/ko-kr';
import ptBR from '@arco-design/web-vue/es/locale/lang/pt-pt';
import thTH from '@arco-design/web-vue/es/locale/lang/th-th';
import zhCN from '@arco-design/web-vue/es/locale/lang/zh-cn';
import zhTW from '@arco-design/web-vue/es/locale/lang/zh-tw';



import {
    dateEnUS,
    dateDeDE,
    dateEsAR,
    dateFrFR,
    dateIdID,
    dateItIT,
    dateJaJP,
    dateKoKR,
    datePtBR,
    dateThTH,
    dateZhCN,
    dateZhTW,

} from 'naive-ui'

import {useStore} from "@/store/index.js";
import {useNavigatorLanguage} from "@vueuse/core";

export const defaultLocale = "en-us"

export const localeList = {
    'zh-cn': {name:"简体中文",lang:zh_CN, arcoLang:zhCN, naiveDateLang: dateZhCN},
    'zh-tw': {name:"繁體中文",lang:zh_TW, arcoLang:zhTW, naiveDateLang: dateZhTW},
    'en-us': {name:"English",lang:en_US, arcoLang:enUS, naiveDateLang: dateEnUS},
    'de-de': {name:"Deutsch",lang:de_DE, arcoLang:deDE, naiveDateLang: dateDeDE},
    'es-es': {name:"Español",lang:es_ES, arcoLang:esES, naiveDateLang: dateEsAR},
    'fr-fr': {name:"Français",lang:fr_FR, arcoLang:frFR, naiveDateLang: dateFrFR},
    'id-id': {name:"Indonesian",lang:id_ID, arcoLang:idID, naiveDateLang: dateIdID},
    'it-it': {name:"Italiano",lang:it_IT, arcoLang:itIT, naiveDateLang: dateItIT},
    'ja-jp': {name:"日本語",lang:ja_JP, arcoLang:jaJP, naiveDateLang: dateJaJP},
    'ko-kr': {name:"한국어",lang:ko_KR, arcoLang:koOR, naiveDateLang: dateKoKR},
    'pt-br': {name:"Português",lang:pt_BR, arcoLang:ptBR, naiveDateLang: datePtBR},
    'th-th': {name:"ภาษาไทย",lang:th_TH, arcoLang:thTH, naiveDateLang: dateThTH},
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

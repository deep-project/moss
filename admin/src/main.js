import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import {createPinia} from "pinia";
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import router from './router'
import { i18n } from './locale'

import "tailwindcss/tailwind.css"
import '@arco-design/web-vue/dist/arco.less';
import './style.css'

createApp(App).use(createPinia()).use(router).use(i18n).use(ArcoVue).use(ArcoVueIcon).mount('#app')

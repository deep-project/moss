import { createRouter, createWebHistory } from 'vue-router'
import routes from './routes'
import {useTitle} from '@vueuse/core'
import {t} from "@/locale";
import {computed} from 'vue'

const router = createRouter({
    history: createWebHistory(process.env.NODE_ENV === "production" ? "/{{__DIR__}}" : "/"),
    routes
})


router.beforeEach((to, from, next) => {
    useTitle(computed(()=>t(to.meta.title || to.name.split("-")[0])))
    next()
})

export default router
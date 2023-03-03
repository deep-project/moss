


import layoutMain from '@/layout/main/index.vue'
import { IconNav,IconDashboard,IconLayers,IconFile,IconTag,IconRelation,IconComputer,IconSettings,IconStorage } from '@arco-design/web-vue/es/icon';

export default [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/admin/login.vue')
    },
    {
        path: '/createAdmin',
        name: 'createAdmin',
        component: () => import('@/views/admin/createAdmin.vue')
    },
    {
        path: '/',
        name:"index",
        redirect: '/dashboard',
        component: layoutMain,
        children: [
            {
                path: '/dashboard',
                name: 'dashboard',
                meta: { icon:IconDashboard },
                component: () => import('@/views/dashboard/index.vue'),
            },
            {
                path: '/articles',
                name: 'articles',
                meta: {icon:IconFile},
                component: () => import('@/views/article/index.vue'),
            },
            {
                path: '/categories',
                name: 'categories',
                meta: {icon:IconNav},
                component:() => import('@/views/category/index.vue'),
            },
            {
                path: '/tags',
                name: 'tags',
                meta: {icon:IconTag},
                component: () => import('@/views/tag/index.vue'),
            },
            {
                path: '/links',
                name: 'links',
                meta: {icon:IconRelation},
                component: () => import('@/views/link/index.vue'),
            },
            {
                path: '/storehouse',
                name: 'storehouse',
                meta: {icon:IconStorage},
                component: () => import('@/views/store/index.vue'),
            },
            {
                path: '/logs',
                name: 'logs',
                redirect: { name: 'logs-item', params:{id:''} },
                meta: {icon:IconComputer},
                component: () => import('@/views/log/layout.vue'),
                children: [
                    {
                        path: '/log/:id?',
                        name: 'logs-item',
                        meta:{ lang:'log'},
                        component: () => import('@/views/log/index.vue'),
                    }
                ]
            },
            {
                path: '/plugins',
                name: 'plugins',
                meta: {icon:IconLayers},
                component: () => import('@/views/plugin/index.vue'),
            },
            {
                path: '/configure',
                name: 'config',
                redirect: { name: 'config-item', params:{id:'site'} },
                meta: {icon:IconSettings},
                component: () => import('@/views/config/layout.vue'),
                children: [
                    {
                        path: '/config/:id',
                        name: 'config-item',
                        meta:{ lang:'config'},
                        component: () => import('@/views/config/index.vue'),
                    }
                ]
            },
        ]
    }
]
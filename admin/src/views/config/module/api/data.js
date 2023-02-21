import {t as $t} from "@/locale";
import {computed} from 'vue'

const getListPayLoad = [
    {field:'select',type:'string', description:'Select Fields'},
    {field:'limit',type:'int',description:'Limit'},
    {field:'page',type:'int',description:'Page Number'},
    {field:'order',type:'string',description:'Order'},
]

// article
export const article = computed(()=>[
    { title: 'Create', method: 'POST', url: '/article/create', payload:articlePayload('create')},
    { title: 'Update', method: 'POST', url: '/article/update', payload:articlePayload('update')},
    { title: 'Get', method: 'GET', url: '/article/get/:id'},
    { title: 'List', method: 'POST', url: '/article/list', payload:getListPayLoad},
    { title: 'Delete', method: 'POST', url: '/article/delete/:id'},
    { title: 'Exists Slug', method: 'POST', url: '/article/existsSlug'},
    { title: 'Exists Title', method: 'POST', url: '/article/existsTitle'},
])

function articlePayload(t){
    let res = [
        {field:'title',type:'string',required:true,description:$t('title')},
        {field:"content",type:'string',required:true,description:$t('content')},
        {field:'slug',type:'string',required:t==='update',description:$t('slug')},
        {field:'category_id',type:'int',description:$t('category_id')},
        {field:'create_time',type:'timestamp',description:$t('createTime')},
        {field:'views',type:'int',description:$t('views')},
        {field:'thumbnail',type:'string',description:$t('thumbnail')},
        {field:'keywords',type:'string',description:$t('keywords')},
        {field:'description',type:'string',description:$t('description')},
        {field:'tags',type: '[]string',description: $t('tags'), example:"['tag1','tag2','tag3']"},
        {field:'category_name',type: 'string',description: $t('category') + ' (If the category id is undefined)'},
        {field:'extends',type:'json',description:$t('extends'), example:"{'source':'google.com','author':'Lucy'}"},
    ]
    if(t==='update') res.unshift({field:'id',type:'int',required:true,description:$t('id')})
    return res
}

// category
export const category = computed(()=>[
    { title: 'Create', method: 'POST', url: '/category/create' ,payload:categoryPayload()},
    { title: 'Update', method: 'POST', url: '/category/update' ,payload:categoryPayload('update')},
    { title: 'Get', method: 'GET', url: '/category/get/:id'},
    { title: 'List', method: 'POST', url: '/category/list', payload:getListPayLoad},
    { title: 'Delete', method: 'POST', url: '/category/delete/:id'},
    { title: 'Exists Slug', method: 'POST', url: '/category/existsSlug'},
    { title: 'Exists Name', method: 'POST', url: '/category/existsName'},
])

function categoryPayload(t){
    let res = [
        {field:'name',type:'string', required:true, description:$t('name')},
        {field:'slug',type:'string',required:t==='update',description:$t('slug')},
        {field:'parent_id',type:'int',description:'parent id'},
        {field:'create_time',type:'timestamp',description:$t('createTime')},
        {field:'title',type:'string',description:$t('title')},
        {field:'keywords',type:'string',description:$t('keywords')},
        {field:'description',type:'string',description:$t('description')},
    ]
    if(t==='update') res.unshift({field:'id',type:'int',required:true,description:$t('id')})
    return res
}

// tag
export const tag = computed(()=>[
    { title: 'Create', method: 'POST', url: '/tag/create', payload:tagPayload()},
    { title: 'Update', method: 'POST', url: '/tag/update', payload:tagPayload('update')},
    { title: 'Get', method: 'GET', url: '/tag/get/:id'},
    { title: 'List', method: 'POST', url: '/tag/list', payload:getListPayLoad},
    { title: 'Delete', method: 'POST', url: '/tag/delete/:id', examples:'POST /data/tag/delete/2'},
    { title: 'Exists Slug', method: 'POST', url: '/tag/existsSlug'},
    { title: 'Exists Name', method: 'POST', url: '/tag/existsName'},
])

function tagPayload(t){
    let res = [
        {field:'name',type:'string', required:true, description:$t('name')},
        {field:'slug',type:'string',required:t==='update',description:$t('slug')},
        {field:'create_time',type:'timestamp',description:$t('createTime')},
        {field:'title',type:'string',description:$t('title')},
        {field:'keywords',type:'string',description:$t('keywords')},
        {field:'description',type:'string',description:$t('description')},
    ]
    if(t==='update') res.unshift({field:'id',type:'int',required:true,description:$t('id')})
    return res
}

// link
export const link = computed(()=>[
    {title: 'Create', method: 'POST', url: '/link/create', payload:linkPayload()},
    {title: 'Update', method: 'POST', url: '/link/update', payload:linkPayload('update')},
    {title: 'Get', method: 'GET', url: '/link/get/:id'},
    {title: 'List', method: 'POST', url: '/link/list', payload:getListPayLoad},
    {title: 'Delete', method: 'POST', url: '/link/delete/:id'},
    {title: 'Exists URL', method: 'POST', url: '/link/existsURL'},
])

function linkPayload(t){
    let res = [
        {field:'name',type:'string', required:true, description:$t('name')},
        {field:'url',type:'string', required:true, description:$t('url')},
        {field:'status',type:'bool',description:$t('status')},
        {field:'logo',type:'string',description:'logo'},
        {field:'note',type:'string',description:$t('note')},
        {field:'create_time',type:'timestamp',description:$t('createTime')},
        {field:'expire_time',type:'timestamp',description:$t('expireTime')},
        {field:'detect',type:'bool',description:$t('detect')},
        {field:'detect_delay',type:'int',description:$t('detect') + ' ' + $t('delay') + '('+$t('minutes')+')'},
    ]
    if(t==='update') res.unshift({field:'id',type:'int',required:true,description:$t('id')})
    return res
}


// store
export const store = computed(()=>[
    {title: 'Create', method: 'POST', url: '/store/create', payload:storePayload()},
    {title: 'Update', method: 'POST', url: '/store/update', payload:storePayload('update')},
    {title: 'Get', method: 'GET', url: '/store/get/:id'},
    {title: 'List', method: 'POST', url: '/store/list', payload:getListPayLoad},
    {title: 'Delete', method: 'POST', url: '/store/delete/:id'},
    {title: $t('publish'), method: 'POST', url: '/store/post/:id'},
])

function storePayload(t){
    let res = [
        {field:'title',type:'string', required:true, description:$t('title')},
        {field:'content',type:'string', required:true, description:$t('content')},
        {field:'category_id',type:'int',description:'category id'},
        {field:'category_name',type: 'string',description: $t('category') + ' (If the category id is undefined)'},
        {field:'tags',type: '[]string',description: $t('tags'), example:"['tag1','tag2','tag3']"},
        {field:'slug',type:'string',description:$t('slug')},
        {field:'thumbnail',type:'string',description:$t('thumbnail')},
        {field:'views',type:'int',description:$t('views')},
        {field:'create_time',type:'timestamp',description:$t('createTime')},
        {field:'description',type:'string',description:$t('description')},
        {field:'keywords',type:'string',description:$t('keywords')},
        {field:'extends',type:'json',description:$t('extends'), example:"{'source':'google.com','author':'Lucy'}"},
    ]
    if(t==='update') res.unshift({field:'id',type:'int',required:true,description:$t('id')})
    return res
}
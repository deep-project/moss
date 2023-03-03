import {useGet,usePost,usePostData,useGetData} from "./hook";

export const adminExists = () => useGet("/admin/exists");
export const adminCreate = (data) => usePost("/admin/create",data);
export const adminCaptcha = () => useGet("/admin/captcha");
export const adminLogin = (data) => usePost("/admin/login", data);


export const tableList = (modelName, data) => usePostData(`/${modelName}/list`, data);
export const tableCount = (modelName,data) => usePostData(`/${modelName}/count`, data);
export const tableDelete = (modelName,id) => usePost(`/${modelName}/delete/${id}`);
export const tableBatchDelete = (modelName,data) => usePost(`/${modelName}/batchDelete`, data);
export const tableGet = (modelName,id) => useGetData(`/${modelName}/get/${id}`);
export const tableCreate = (modelName,data) => usePost(`/${modelName}/create`,data);
export const tableUpdate = (modelName,data) => usePost(`/${modelName}/update`,data);


export const articleCreateTagByNameList = (articleID, nameList) => usePost(`/article/createTagByNameList/${articleID}`, nameList);
export const articleDeleteTagByIds = (articleID, tagIds) => usePost(`/article/deleteTagByIds/${articleID}`, tagIds);
export const articleBatchSetCategory = (category_id,data) => usePost(`/article/batchSetCategory/${category_id}`,data);

export const categoryTree = ()=> useGetData('/category/tree')
export const categoryBatchSetParentCategory = (parent_id,data)=> usePost(`/category/batchSetParentCategory/${parent_id}`,data)

export const tagsByArticleID = (id) => useGetData(`/tag/list/article/${id}`);
export const tagGetByIds = (ids) => usePostData(`/tag/getByIds`, ids);
export const linkStatus = (id,status) => usePostData(`/link/status/${id}`, {status:status});
export const storePost = (id) => usePost(`/store/post/${id}`);






export const logInit = () => usePost(`/log/init`);
export const logRead = (id, params) => useGetData(`/log/read/${id}`,params);

export const configList = () => useGet("/config");
export const configGet = (id) => useGet(`/config/${id}`);
export const configPost = (id, data) => usePost(`/config/${id}`, data);

export const uploadInit = () => usePost(`/upload/init`);
export const upload = (f) => usePost(`/upload`,f);

export const themeInit = () => usePost(`/theme/init`);
export const themeList = () => useGetData(`/theme/list`);
export const themeScreenshot = (id) => useGetData(`/theme/screenshot/${id}`);

export const cacheInit = () => usePost(`/cache/init`);
export const cacheClear = (name) => usePost(`/cache/clear/${name}`);

export const routerReload = () => usePost(`/router/reload`);

export const pluginList = () => useGetData(`/plugin/list`);
export const pluginOptions = (id) => useGetData(`/plugin/options/${id}`);
export const pluginSaveOptions = (id,data) => usePost(`/plugin/saveOptions/${id}`,data);
export const pluginRun = (id) => usePost(`/plugin/run/${id}`);
export const pluginCronStart= (id) => usePost(`/plugin/cron/start/${id}`);
export const pluginCronStop = (id) => usePost(`/plugin/cron/stop/${id}`);
export const pluginCronExp = (id,data) => usePost(`/plugin/cron/exp/${id}`, data);
export const pluginLogList = (id,params) => useGetData(`/plugin/log/list/${id}`,params);


export const dashboardData = (id) => useGetData(`/dashboard/${id}`);
import axios from './axios'

export async function useGet(url, params){
    const response = await axios.get(url, {params: params})
    return response.data
}

export async function usePost(url, data){
    const response = await axios.post(url, data)
    return response.data
}

export async function useGetData(url, params){
    const response =  await axios.get(url, {params: params})
    return response.data.data
}


export async function usePostData(url, data){
    const response =  await axios.post(url, data)
    return response.data.data
}

export function useGetRaw(url, params){
    return axios.get(url, {params: params})
}

export function usePostRaw(url, data){
    return axios.post(url, data)
}
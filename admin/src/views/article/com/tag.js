


export function getIndexByName(list, name){
    for(let i in list){
        if(list[i].name === name) return i
    }
    return -1
}

export function existName(list, name){
    return getIndexByName(list, name) > -1
}


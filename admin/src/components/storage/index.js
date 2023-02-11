

let all = [
    {value:"local",label:"local file"},
    {value:"ftp",label:"ftp"},
    {value:"s3",label:"amazon s3"},
    {value:"b2",label:"backblaze b2"},
    {value:"cos",label:"tencentcloud cos"},
    {value:"oss",label:"alibabacloud oss"},
    {value:"badger",label:"badgerDB"},
    {value:"redis",label:"redis"},
    {value:"memcached",label:"memcached"},
]

export function getOptions(allowed=[]){
    if(allowed.length === 0){
        return all
    }
    let res = []
    for (let name of allowed){
        for (let a of all){
            if(name === a.value){
                res.push(a)
            }
        }
    }
    return res
}
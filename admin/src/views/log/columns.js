import {useDeepCopy} from "@/hooks/utils.js";

export const columnAPP = [
    {
        width: 30,
        slotName:'detail'
    },
    {
        title: 'time',
        dataIndex: 'time',
        slotName:'time',
        width: 120,
    },
    {
        title: 'level',
        dataIndex: 'level',
        slotName:'level',
        width: 100,
    },
    {
        title: 'file',
        dataIndex: 'file',
        ellipsis:true,
        width: 200,
    },
    {
        title: 'message',
        dataIndex: 'msg',
        ellipsis:true,
        width: 200,
    },
    {
        title: 'more',
        dataIndex: 'more',
        slotName:'more',
        ellipsis:true,
        width: 300,
    },
]

export const columnSQL = [
    {
        width: 30,
        slotName:'detail'
    },
    {
        title: 'time',
        dataIndex: 'time',
        slotName:'time',
        width: 120,
    },
    {
        title: 'take',
        dataIndex: 'take',
        slotName:'take',
        width: 100,
    },
    {
        title: 'file',
        dataIndex: 'file',
        ellipsis:true,
        width: 170,
    },
    {
        title: 'rows',
        dataIndex: 'rows',
        slotName:'rows',
        width: 80,
    },
    {
        title: 'sql',
        dataIndex: 'sql',
        ellipsis:true,
        width: 400,
    },
]

export const columnHTTP = [
    {
        width: 30,
        slotName:'detail'
    },
    {
        title: 'time',
        dataIndex: 'time',
        slotName:'time',
        width: 120,
    },
    {
        title: 'take',
        dataIndex: 'take',
        slotName:'take',
        width: 100,
    },
    {
        title: 'status',
        dataIndex: 'status',
        slotName:'status',
        width: 80,
    },
    {
        title: 'depth',
        dataIndex: 'depth',
        width: 70,
    },
    {
        title: 'ip',
        dataIndex: 'ip',
        width: 140,
        ellipsis:true,
        tooltip:true,
    },
    {
        title: 'region',
        dataIndex: 'region',
        width: 220,
        ellipsis:true,
        tooltip:true,
    },
    {
        title: 'method',
        dataIndex: 'method',
        width: 80,
    },
    {
        title: 'url',
        dataIndex: 'url',
        width: 300,
        ellipsis:true,
        tooltip:true,
        slotName:'url',
    },
    {
        title: 'referer',
        dataIndex: 'referer',
        width: 300,
        ellipsis:true,
        tooltip:true,
        slotName:'url',
    },
    {
        title: 'userAgent',
        dataIndex: 'userAgent',
        width: 1000,
        ellipsis:true,
        tooltip:true,
    },

]

export function columnSpider(){
    let col = useDeepCopy(columnHTTP)
    col.splice(2, 0, {
        title: 'feature',
        dataIndex: 'feature',
        width: 130,
        ellipsis:true,
        tooltip:true,
    });
    return col
}
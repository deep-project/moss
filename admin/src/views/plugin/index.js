

export const columns = [
    {
        title: 'ID',
        dataIndex: 'id',
        slotName:'id',
        width:240,
        ellipsis:true,
    },
    {
        slotName: 'action',
        width:170,
    },
    {
        title: 'Cron',
        slotName: 'cron',
        width:240,
    },
    {
        title: 'Last run time',
        dataIndex: 'run_time',
        slotName: 'time',
        width:140,
    },
    {
        title: 'Next run time',
        dataIndex: 'next_run_time',
        slotName: 'time',
        width:140,
    },
    {
        title: 'Run duration',
        slotName: 'runDuration',
        width:140,
    },
    {
        title: 'Run count',
        dataIndex: 'run_count',
        slotName: 'runCount',
        width:140,
    },
    {
        title: 'Run error',
        dataIndex: 'run_error',
    },
]
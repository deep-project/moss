<template>
  <a-descriptions :column="1" bordered>
    <a-descriptions-item label="time"><n-time :time="data.time" /></a-descriptions-item>
    <a-descriptions-item label="level">{{data.level}}</a-descriptions-item>
    <a-descriptions-item label="take">{{data.take}} ms</a-descriptions-item>
    <a-descriptions-item label="file">{{data.file}}</a-descriptions-item>
    <a-descriptions-item label="rows">{{data.rows}}</a-descriptions-item>
    <a-descriptions-item label="msg">{{data.msg}}</a-descriptions-item>
    <a-descriptions-item label="sql">
      <div class="py-2 my-2">
        <code class="sql font-mono whitespace-pre-line" v-html="prettify(data.sql)"></code>
      </div>
    </a-descriptions-item>
  </a-descriptions>
</template>

<script setup>
  import { NTime } from 'naive-ui'

  defineProps({data:Object})

  let sqlKeywords = [
    'select ','from ','distinct',' where ',' order by ',' group by ',' insert into',' inner join',' left join',' and ',
    'having', 'values', 'on conflict', 'do update set', ' create table',' limit ',' as ',' update ',' offset ',' set ','do nothing',
    'alter table ','alter column ','type ','default ','in '
  ]
  function prettify(sql){
    if(!sql) return sql
    for(let i=0;i<sqlKeywords.length;i++){
      let v = sqlKeywords[i]
      let v2 = v.toUpperCase()
      sql = sql.replaceAll(v, v2)
      sql = sql.replaceAll(v2,'<b>'+v2+'</b>')
    }
    return sql.replace(/"([^"]*)"/g, `"<i>$1</i>"`).replace(/'([^']*)'/g, `'<i>$1</i>'`)
  }

</script>

<style>
.sql>b{
  font-weight: normal;
  color:#FF66CC;
}
.sql>i{
  color:#999;
  margin:0 3px;
}
</style>
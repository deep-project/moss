package ip2region

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"moss/resources"
	"strings"
)

var Client *xdb.Searcher

// db下载列表
// https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.xdb
// https://fastly.jsdelivr.net/gh/lionsoul2014/ip2region@master/data/ip2region.xdb
// https://cdn.jsdelivr.net/gh/lionsoul2014/ip2region@master/data/ip2region.xdb

func init() {
	if err := Init(); err != nil {
		fmt.Println("init ip2region failed", err.Error())
	}
}

func Init() error {
	b, err := resources.App.ReadFile("app/ip2region.xdb")
	if err != nil {
		return err
	}
	searcher, err := xdb.NewWithBuffer(b)
	if err != nil {
		return err
	}
	Client = searcher
	return nil
}

func Region(ip string) string {
	if Client == nil {
		return ""
	}
	if ip == "" {
		return ""
	}
	res, _ := Client.SearchByStr(ip)
	return format(res)
}

func format(re string) string {
	arr := strings.Split(re, "|")
	var newArr []string
	for _, v := range arr {
		if v != "0" && v != "" && v != " " {
			newArr = append(newArr, v)
		}
	}
	return strings.Join(newArr, " ")
}

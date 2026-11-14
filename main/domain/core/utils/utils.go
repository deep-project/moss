package utils

import (
	"moss/domain/core/aggregate"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// SortByIds 根据ids排序
func SortByIds[M aggregate.EntityInterface](items []M, ids []int) (res []M) {
	if len(items) < 2 || len(ids) < 2 {
		return items
	}
	m := SliceToMap(items, "ID")
	for _, id := range ids {
		item, ok := m[id]
		if !ok {
			continue
		}
		res = append(res, item)
	}
	return
}

// SliceToMap 切片转map
func SliceToMap[M aggregate.EntityInterface](s []M, field string) map[any]M {
	var res = make(map[any]M)
	for _, v := range s {
		id := int(reflect.ValueOf(v).FieldByName(field).Int())
		res[id] = v
	}
	return res
}

func TodayBeginTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

func ValidateOrderInput(val any) any {
	str, ok := val.(string)
	if !ok || str == "" {
		return val
	}
	str = strings.ToLower(strings.TrimSpace(str))
	if str == "rand()" || str == "random()" {
		return str
	}
	if regexp.MustCompile(`^([a-z0-9_]+(?:\s+(asc|desc))?\s*,\s*)*[a-z0-9_]+(?:\s+(asc|desc))?$`).MatchString(str) {
		return str
	}
	if regexp.MustCompile(`^(count|sum|max|min|avg|length|year)\((\w+)\)(?:\s+(asc|desc))?$`).MatchString(str) {
		return str
	}
	return nil
}

// ValidateSelectInput 校验 select 输入，仅允许逗号分隔的列名（可带表前缀或 *）。
// 其他内容（空格分隔的关键字、括号、子查询、引号等）一律丢弃，
// 返回空字符串表示使用默认全部字段，防止通过 select 子句进行 SQL 注入。
func ValidateSelectInput(val string) string {
	val = strings.TrimSpace(val)
	if val == "" {
		return ""
	}
	if regexp.MustCompile(`^(\*|[a-zA-Z0-9_]+(\.[a-zA-Z0-9_]+)?)(\s*,\s*(\*|[a-zA-Z0-9_]+(\.[a-zA-Z0-9_]+)?))*$`).MatchString(val) {
		return val
	}
	return ""
}

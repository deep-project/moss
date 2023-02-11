package utils

import (
	"github.com/duke-git/lancet/v2/random"
	"time"
)

type Utils struct{}

func New() *Utils {
	return &Utils{}
}

func (*Utils) Pagination(page int, pageTotal int, limit int) PaginationModel {
	return Pagination(page, pageTotal, limit)
}

func (*Utils) FormatTimestamp(ts int64, layout string) string {
	return FormatTimestamp(ts, layout)
}

func (*Utils) RandString(length int) string {
	return random.RandString(length)
}

func (*Utils) RandInt(min int, max int) int {
	return random.RandInt(min, max)
}

func (*Utils) UUID() (string, error) {
	return random.UUIdV4()
}

func FormatTimestamp(ts int64, layout string) string {
	if ts == 0 {
		return ""
	} else if ts == -1 {
		return time.Now().Format(layout)
	}
	return time.Unix(ts, 0).Format(layout)
}

type PaginationModel struct {
	Begin int
	End   int
}

func Pagination(page, pageTotal, limit int) PaginationModel {
	if limit <= 0 {
		return PaginationModel{Begin: 0, End: 0}
	}
	if pageTotal < limit {
		limit = pageTotal
	}
	var begin = page - (limit / 2)
	if begin <= 0 {
		begin = 1
	}
	var end = limit + begin - 1
	if end > pageTotal {
		begin = begin - (end - pageTotal)
		end = pageTotal
		if begin < 1 {
			begin = 1
		}
	}
	return PaginationModel{Begin: begin, End: end}
}

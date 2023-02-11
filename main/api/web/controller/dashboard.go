package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"moss/api/web/mapper"
	appService "moss/application/service"
	"moss/domain/core/service"
	"moss/infrastructure/persistent/db"
	"time"
)

var Dashboard = new(dashboard)

type dashboard struct {
}

func (d *dashboard) Controller(ctx *fiber.Ctx) (err error) {
	var data any
	switch ctx.Params("id") {
	case "systemLoad":
		info, _ := load.Avg()
		data = info.Load5
	case "systemCPU":
		v, _ := cpu.Percent(time.Second, false)
		data = v[0]
	case "systemMemory":
		v, _ := mem.VirtualMemory()
		data = v.UsedPercent
	case "systemDisk":
		parts, _ := disk.Partitions(true)
		for _, part := range parts {
			diskInfo, _ := disk.Usage(part.Mountpoint)
			data = diskInfo.UsedPercent
			break
		}
	case "database":
		data = db.GetSize()
	case "log":
		data, err = appService.LogDirSize()
	case "cache":
		data, err = appService.CacheSize()
	case "articleTotal":
		data, err = service.Article.CountTotal()
	case "articleToday":
		data, err = service.Article.CountToday()
	case "articleYesterday":
		data, err = service.Article.CountYesterday()
	case "categoryTotal":
		data, err = service.Category.CountTotal()
	case "tagTotal":
		data, err = service.Tag.CountTotal()
	case "linkTotal":
		data, err = service.Link.CountTotal()
	default:
		return ctx.JSON(mapper.MessageResult(errors.New("id is undefined")))
	}
	return ctx.JSON(mapper.MessageResultData(data, err))
}

package service

import (
	"bytes"
	"errors"
	"github.com/google/pprof/profile"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"runtime/pprof"
	"time"
)

// SystemLoadPercent 系统平均负载
func SystemLoadPercent() float64 {
	info, _ := load.Avg()
	if info == nil {
		return -1
	}
	return (info.Load1 + info.Load5 + info.Load15) / 3
}

// SystemCPUPercent 系统CPU使用率
func SystemCPUPercent(interval time.Duration) (_ float64, err error) {
	v, err := cpu.Percent(interval, false)
	if err != nil {
		return
	}
	return v[0], nil
}

// SystemMemoryPercent 系统内存使用率
func SystemMemoryPercent() (_ float64, err error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	return v.UsedPercent, nil
}

// SystemDiskPercents 系统硬盘占用率
func SystemDiskPercents() (res []float64, err error) {
	parts, err := disk.Partitions(false)
	if err != nil {
		return
	}
	for _, part := range parts {
		diskInfo, _ := disk.Usage(part.Mountpoint)
		res = append(res, diskInfo.UsedPercent)
	}
	return
}

// AppCPUPercent 应用cpu使用率
func AppCPUPercent() (res float64, err error) {
	p, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return
	}
	if res, err = p.CPUPercent(); err != nil {
		return
	}
	if res > 100 {
		res = 100
	}
	return
}

// AppUsedMemory 应用内存占用（字节）
// 使用读取 profile 的方式统计内存占用，相对比较准确
func AppUsedMemory() (res uint64, err error) {
	// 获取内存 profile 数据
	memProfile := pprof.Lookup("heap")
	if memProfile == nil {
		return res, errors.New("heap profile not found")
	}
	buf := &bytes.Buffer{}
	if err = memProfile.WriteTo(buf, 0); err != nil {
		return
	}
	// 解析缓冲区中的内存 profile 数据
	prof, err := profile.Parse(buf)
	if err != nil {
		return
	}
	var memInuseSpace int64
	for _, sample := range prof.Sample {
		if len(sample.Value) >= 4 {
			memInuseSpace += sample.Value[3]
		}
	}
	return uint64(memInuseSpace), nil
}

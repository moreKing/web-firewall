package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"

	"server/internal/model"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/system/v1"
)

func (c *ControllerV1) GetSystemStatus(ctx context.Context, req *v1.GetSystemStatusReq) (res *v1.GetSystemStatusRes, err error) {

	// cpu数量
	cpuTotal, err := cpu.Counts(false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取cpu信息，请联系技术支持")
	}

	cpuLogicalTotal, err := cpu.Counts(true)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取cpu信息，请联系技术支持")
	}

	//5s cpu使用率 返回一个值
	cpuPercent, err := cpu.Percent(time.Second*1, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取cpu信息，请联系技术支持")
	}

	// 内存
	memory, err := mem.VirtualMemory()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取内存信息，请联系技术支持")
	}
	//memory.
	//memory.Total memory.Free memory.UsedPercent

	partitions, err := disk.Partitions(true)
	var usages []*disk.UsageStat
	for _, partition := range partitions {

		us, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取磁盘信息，请联系技术支持")
		}
		if us.Fstype == "tmpfs" {
			continue
		}
		if us.Total == 0 {
			continue
		}
		usages = append(usages, us)
	}

	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.NewCode(gcode.CodeInvalidOperation, "无法获取磁盘信息，请联系技术支持")
	}

	return &v1.GetSystemStatusRes{
		SystemStatus: &model.SystemStatus{
			CPUTotal:          cpuTotal,
			CpuLogicalTotal:   cpuLogicalTotal,
			CpuPercent:        cpuPercent[0],
			MemoryTotal:       memory.Total,
			MemoryUsed:        memory.Used,
			MemoryUsedPercent: memory.UsedPercent,
			Partitions:        usages,
		},
	}, nil
}

package agent

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"time"
)

type HostInfoFetcher interface {
	Fetch() (AgentInfo, error)
}

type DefaultHostInfoFetcher struct {
}

func (fetcher DefaultHostInfoFetcher) Fetch() (AgentInfo, error) {
	hostInfo, _ := fetcher.fetchHostInfo()
	cpuInfo, _ := fetcher.fetchCpuInfo()
	memInfo, _ := fetcher.fetchMemoryInfo()
	disksInfo, _ := fetcher.fetchDisksInfo()
	return AgentInfo{
		Host:   hostInfo,
		Cpus:   cpuInfo,
		Memory: memInfo,
		Disks:  disksInfo,
	}, nil
}

func (fetcher DefaultHostInfoFetcher) fetchHostInfo() (HostInfo, error) {
	info, err := host.Info()

	if err != nil {
		return HostInfo{}, err
	}

	return HostInfo{
		Hostname:        info.Hostname,
		UpTime:          info.Uptime,
		OS:              info.OS,
		Platform:        info.Platform,
		PlatformFamily:  info.PlatformFamily,
		PlatformVersion: info.PlatformVersion,
	}, nil
}

func (fetcher DefaultHostInfoFetcher) fetchCpuInfo() ([]CpuInfo, error) {
	info, err := cpu.Info()

	if err != nil {
		return []CpuInfo{}, err
	}

	usages, err := cpu.Percent(1*time.Second, true)

	if err != nil {
		usages = make([]float64, 0)
	}

	totalUsage, err := cpu.Percent(1*time.Second, false)

	if err != nil {
		totalUsage = make([]float64, 1)
	}

	perCpuInfo := make([]CpuInfo, 0)

	for _, cpuInfo := range info {
		if cpuInfo.Cores != 0 {
			perCpuInfo = append(perCpuInfo, CpuInfo{
				Vendor:            cpuInfo.VendorID,
				Family:            cpuInfo.Family,
				Model:             cpuInfo.Model,
				Cores:             cpuInfo.Cores,
				ModelName:         cpuInfo.ModelName,
				Frequency:         cpuInfo.Mhz,
				UsagePercent:      usages,
				UsagePercentTotal: totalUsage[0],
			})
		}
	}

	return perCpuInfo, nil
}

func (fetcher DefaultHostInfoFetcher) fetchMemoryInfo() (Memory, error) {

	stats, err := mem.VirtualMemory()

	if err != nil {
		return Memory{}, err
	}

	return Memory{
		Total:       stats.Total,
		Used:        stats.Used,
		Free:        stats.Free,
		UsedPercent: stats.UsedPercent,
	}, nil

}

func (fetcher DefaultHostInfoFetcher) fetchDisksInfo() ([]DiskInfo, error) {
	partitions, err := disk.Partitions(false)

	if err != nil {
		return []DiskInfo{}, err
	}

	diskStats := make([]DiskInfo, 0)

	for _, partition := range partitions {

		info, err := disk.Usage(partition.Mountpoint)

		if err == nil {
			if info.Total != 0 {
				diskStats = append(diskStats, DiskInfo{
					Name:              partition.Device,
					Mount:             partition.Mountpoint,
					Total:             info.Total,
					Free:              info.Free,
					Used:              info.Used,
					UsedPercent:       info.UsedPercent,
					InodesTotal:       info.InodesTotal,
					InodesUsed:        info.InodesUsed,
					InodesFree:        info.InodesFree,
					InodesUsedPercent: info.InodesUsedPercent,
				})
			}
		}
	}

	return diskStats, nil
}

package metrics

import (
	"github.com/cactus/go-statsd-client/statsd"
	"time"
	"github.com/GaruGaru/Warden/agent"
	"strings"
	"strconv"
)

func NewStatsdMetricsReporter(address string, prefix string) (StatsdMetricsReporter, error) {
	client, err := statsd.NewBufferedClient(address, prefix, 500*time.Millisecond, 0)
	if err != nil {
		return StatsdMetricsReporter{}, err
	}

	return StatsdMetricsReporter{
		Client: client,
	}, nil
}

type StatsdMetricsReporter struct {
	Client statsd.Statter
}

func (reporter StatsdMetricsReporter) Send(info agent.AgentInfo) error {

	reporter.sendCpuInfo(info.CpusInfo)
	reporter.sendMemoryInfo(info.MemoryInfo)
	reporter.sendDisksInfo(info.Disks)

	return nil
}

func (reporter StatsdMetricsReporter) sendDisksInfo(info []agent.DiskInfo) {
	for _, disk := range info {
		baseKey := key("disk", disk.Mount)

		reporter.Client.Gauge(key(baseKey, "total"), int64(disk.Total), 1)
		reporter.Client.Gauge(key(baseKey, "free"), int64(disk.Free), 1)
		reporter.Client.Gauge(key(baseKey, "used"), int64(disk.Used), 1)
		reporter.Client.Gauge(key(baseKey, "used_percent"), int64(disk.UsedPercent), 1)

		reporter.Client.Gauge(key(baseKey, "inodes_total"), int64(disk.InodesTotal), 1)
		reporter.Client.Gauge(key(baseKey, "inodes_used"), int64(disk.Used), 1)
		reporter.Client.Gauge(key(baseKey, "inodes_free"), int64(disk.InodesFree), 1)
		reporter.Client.Gauge(key(baseKey, "inodes_used_percent"), int64(disk.InodesUsedPercent), 1)

	}
}

func (reporter StatsdMetricsReporter) sendMemoryInfo(info agent.MemoryInfo) {
	baseKey := "memory"
	reporter.Client.Gauge(key(baseKey, "total"), int64(info.Total), 1)
	reporter.Client.Gauge(key(baseKey, "free"), int64(info.Free), 1)
	reporter.Client.Gauge(key(baseKey, "used"), int64(info.Used), 1)
	reporter.Client.Gauge(key(baseKey, "used_percent"), int64(info.UsedPercent), 1)
}

func (reporter StatsdMetricsReporter) sendCpuInfo(info []agent.CpuInfo) {
	for i, cpu := range info {
		baseKey := key("cpu", strconv.Itoa(i))

		reporter.Client.Gauge(key(baseKey, "usage_total"), int64(cpu.UsagePercentTotal), 1)
		reporter.Client.Gauge(key(baseKey, "frequency"), int64(cpu.Frequency), 1)

		for ic, usage := range cpu.UsagePercent {
			coreKey := strings.Join([]string{baseKey, "cores", strconv.Itoa(ic)}, ".")
			reporter.Client.Gauge(key(coreKey, "usage"), int64(usage), 1)
		}

	}
}

func key(base string, key string) string {
	return strings.Join([]string{base, key}, ".")
}

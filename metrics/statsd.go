package metrics

import (
	"github.com/GaruGaru/Warden/agent"
	"github.com/cactus/go-statsd-client/statsd"
	"strconv"
	"strings"
	"time"
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

	nodeName := info.Host.Hostname

	reporter.sendCpuInfo(nodeName, info.CpusInfo)
	reporter.sendMemoryInfo(nodeName, info.MemoryInfo)
	reporter.sendDisksInfo(nodeName, info.Disks)

	return nil
}

func (reporter StatsdMetricsReporter) sendDisksInfo(nodeName string, info []agent.DiskInfo) {
	for _, disk := range info {
		baseKey := key(nodeName, "disk", disk.Mount)

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

func (reporter StatsdMetricsReporter) sendMemoryInfo(nodeName string, info agent.MemoryInfo) {
	baseKey := key(nodeName, "memory")
	reporter.Client.Gauge(key(baseKey, "total"), int64(info.Total), 1)
	reporter.Client.Gauge(key(baseKey, "free"), int64(info.Free), 1)
	reporter.Client.Gauge(key(baseKey, "used"), int64(info.Used), 1)
	reporter.Client.Gauge(key(baseKey, "used_percent"), int64(info.UsedPercent), 1)
}

func (reporter StatsdMetricsReporter) sendCpuInfo(nodeName string, info []agent.CpuInfo) {
	for i, cpu := range info {
		baseKey := key(nodeName, "cpu", strconv.Itoa(i))

		reporter.Client.Gauge(key(baseKey, "usage_total"), int64(cpu.UsagePercentTotal), 1)
		reporter.Client.Gauge(key(baseKey, "frequency"), int64(cpu.Frequency), 1)

		for ic, usage := range cpu.UsagePercent {
			coreKey := strings.Join([]string{baseKey, "cores", strconv.Itoa(ic)}, ".")
			reporter.Client.Gauge(key(coreKey, "usage"), int64(usage), 1)
		}

	}
}

func key(keys ...string) string {
	return strings.Join(keys, ".")
}

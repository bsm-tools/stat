package node

import (
	goNet "net"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Node struct {
	Host      string          `json:"Host"`
	IPAddress []string        `json:"IpAddress"`
	Now       int64           `json:"Now"`
	Runtime   HardwareCollect `json:"Runtime"`
}

// HardwareCollect .
type HardwareCollect struct {
	CPUUsedPercent    float64 `json:"CPUUsedPercent"`
	DiskFree          uint64  `json:"DiskFree"`
	DiskTotal         uint64  `json:"DiskTotal"`
	DiskUsedPercent   float64 `json:"DiskUsedPercent"`
	MemoryFree        uint64  `json:"MemoryFree"`
	MemoryTotal       uint64  `json:"MemoryTotal"`
	MemoryUsedPercent float64 `json:"MemoryUsedPercent"`
	NetIOBytesRecv    uint64  `json:"NetIOBytesRecv"`
	NetIOBytesSent    uint64  `json:"NetIOBytesSent"`
}

func RealTimeHardware() HardwareCollect {
	var collect HardwareCollect
	c, err := cpu.Percent(0, false)
	if err == nil && len(c) >= 0 {
		collect.CPUUsedPercent = c[0]
	}

	v, err := mem.VirtualMemory()
	if err == nil {
		collect.MemoryTotal = v.Total
		collect.MemoryFree = v.Free
		collect.MemoryUsedPercent = v.UsedPercent
	}

	d, err := disk.Usage("/")
	if err == nil {
		collect.DiskTotal = d.Total
		collect.DiskFree = d.Free
		collect.DiskUsedPercent = d.UsedPercent
	}

	// 流量
	n, err := net.IOCounters(false)
	if err == nil && len(c) >= 0 {
		collect.NetIOBytesSent = n[0].BytesSent
		collect.NetIOBytesRecv = n[0].BytesRecv
	}
	return collect
}
func LocalIPv4s() ([]string, error) {

	var ips []string
	addrs, _ := goNet.InterfaceAddrs()

	for _, addr := range addrs {
		if ipnet, ok := addr.(*goNet.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				locIP := ipnet.IP.To4().String()
				if locIP[0:7] != "169.254" {
					ips = append(ips, locIP)
				}
			}
		}
	}

	return ips, nil
}

func Stat() Node {
	name, _ := os.Hostname()
	ips, _ := LocalIPv4s()
	return Node{
		Host:      name,
		IPAddress: ips,
		Runtime:   RealTimeHardware(),
		Now:       time.Now().UnixMicro(),
	}
}

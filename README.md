# stat
response the server running status in JSON format

# Install  
```
curl -LO https://github.com/bsm-tools/stat/releases/download/v1.0/stat-v1.0-linux-amd64  
chmod +x stat-v1.0-linux-amd64  
./stat-v1.0-linux-amd64  
2024/11/24 23:43:44 Starting stat service on port(ENV:STAT_PORT/default:9030) 9030
```

# Import  
```
package main

import (
  "log"
  "github.com/bsm-tools/stat/node"
)

func main(){
  log.Println(node.Stat())
}

```
# Output
```
{
    "Host": "DESKTOP-0U87KOS",
    "IpAddress": [
        "26.26.26.1",
        "192.168.31.35"
    ],
    "Now": 1733418655259594,
    "Runtime": {
        "CPUUsedPercent": 19.6360153256705,
        "DiskFree": 93692014592,
        "DiskTotal": 361259601920,
        "DiskUsedPercent": 74.06518357047078,
        "MemoryFree": 5934952448,
        "MemoryTotal": 17088331776,
        "MemoryUsedPercent": 65,
        "NetIOBytesRecv": 17844934520,
        "NetIOBytesSent": 3215972763
    }
}
```

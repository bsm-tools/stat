# stat
response the server running status in JSON format

# Install  
```
curl -LO https://github.com/yanweidong/stat/releases/download/v1.0/stat-v1.0-linux-amd64  
chmod +x stat-v1.0-linux-amd64  
./stat-v1.0-linux-amd64  
2024/11/24 23:43:44 Starting stat service on port(ENV:STAT_PORT/default:9030) 9030
```

# Import  
```
package main

import (
  "log"
  "github.com/yanweidong/stat/node"
)

func main(){
  log.Println(node.Stat())
}

```

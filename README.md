# go linux tool

Usage
---------------

```go
import (
	"log"

	linuxtool "github.com/vogo/go-linux-tool"
)

stat, err := linuxtool.ReadStat("/proc/stat")
if err != nil {
	log.Fatal("stat read fail")
}

for _, s := range stat.CPUStats {
	// s.User
	// s.Nice
	// s.System
	// s.Idle
	// s.IOWait
}

// stat.CPUStatAll
// stat.CPUStats
// stat.Processes
// stat.BootTime
// ... etc
```

Documentation
---------------

Full documentation is available at [Godoc](https://godoc.org/github.com/vogo/go-linux-tool).


Reference
------------

* http://www.mjmwired.net/kernel/Documentation/filesystems/proc.txt

License
-------

go-linux-tool is distributed under the MIT license.

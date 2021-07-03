# gopython
Let you write go as if you were writing Python.

## Usage
```bash
go get github.com/yingshaoxo/gopython
```

```go
package main

import (
	"log"

	port_scanner "github.com/yingshaoxo/gopython/network"
)

func main() {
	urls := port_scanner.ScanPorts("localhost", 0, 65535)
	log.Println(urls)
}

```
# gopython
Let you write go as if you were writing Python.

## Usage
```bash
go get github.com/yingshaoxo/gopython
```

### try-catch
```go
package main

import (
	"log"
	"fmt"

	error_tool "github.com/yingshaoxo/gopython/error_tool"
)

func main() {
	error_tool.Try(func() {
		i := 3
		for i != -1 {
			result := 100 / i
			fmt.Println(result, i)

			i -= 1
		}
	}).Catch(func(err string) {
		log.Println(err)
	})
}
```

### JWT (json web token)
```go
package main

import (
	"log"

	jwt_tool "github.com/yingshaoxo/gopython/jwt_tool"
)

func main() {
	secret := "no way"

	data := make(map[string]interface{})
	data["user"] = "yingshaoxo"

	jwt_string := jwt_tool.Jwt_encode(data, secret)
	log.Println(jwt_string)

	new_data, _ := jwt_tool.Jwt_decode(jwt_string, secret)

	log.Println(data)
	log.Println(new_data)
}
```

### network
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
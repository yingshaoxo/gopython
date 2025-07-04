# gopython
Let you write go as if you were writing Python.

> **Golang is garbage, you'd better use Python for your safety.**

> **You may want to kill yourself at sometimes if you persistent to use Golang.**

## Usage
```bash
go get github.com/yingshaoxo/gopython
```

## Compile
You know what? You can use golang1.10.8 in both windowsXP and old linux to build linux software.

You just need to create a new folder and put your main go file into it, then inside of that folder, you do following:

```
CGO_ENABLED=0 GOOS=linux GOARCH=386 wine go.exe build -a -ldflags '-extldflags "-static"' -o test.run ./
```

### Try-Catch
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

### Nullable
```go
package main

import (
	"fmt"

	variable_tool "github.com/yingshaoxo/gopython/variable_tool"
)

func main() {
	hi := "hi"
	ok := "ok"

	a_string := variable_tool.Nullable(&hi)
	a_string.Is_null = true
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)

	a_string.Value = &ok
	a_string.Is_null = false
	fmt.Println(*a_string.Value)
	fmt.Println(a_string.Is_null)
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

### Network
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

# Memory status

Retrieve memory status information from the operating system.

Windows, Linux & FreeBSD support so far.

Example:
```go
package main

import "fmt"
import "github.com/howeyc/memstatus"

func main() {
	physical, virtual := memstatus.MemStatus()
    fmt.Println(physical, " ", virtual)
}
```

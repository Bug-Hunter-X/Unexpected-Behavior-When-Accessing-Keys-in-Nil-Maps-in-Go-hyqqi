```go
package main

import "fmt"

func main() {
    var m map[string]int
    if m == nil {
        m = make(map[string]int)
    }
    m["a"] = 10
    fmt.Println(m["a"])
    fmt.Println(m==nil) 

    //Safe way to get value from the map
    value, ok := m["b"]
    if ok {
        fmt.Println("Value of b:", value)
    } else {
        fmt.Println("Key b not found in the map")
    }
}
```
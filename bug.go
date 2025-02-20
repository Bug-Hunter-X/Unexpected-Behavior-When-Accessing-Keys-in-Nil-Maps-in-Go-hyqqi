```go
package main

import "fmt"

func main() {
    var m map[string]int
    fmt.Println(m["a"]) // This will not panic, but return the zero value of int, which is 0.
    fmt.Println(m == nil) //This will print true, showing that m is nil.

    //This will panic if you try to access a non-existent key without checking if the map is nil
    //m["b"] = 10
    //fmt.Println(m["b"])

}
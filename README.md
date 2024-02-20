go-sortedkeys
=============

[![GoDoc](https://godoc.org/github.com/hymkor/go-sortedkeys?status.svg)](https://godoc.org/github.com/hymkor/go-sortedkeys)

Make iterator for key-value pairs of map by sorted order

example
-------

```example.go
package main

import (
    "github.com/hymkor/go-sortedkeys"
)

func main() {
    sample := map[string]string{
        "A": "1",
        "B": "2",
        "C": "3",
        "D": "4",
    }

    println("*** for key,val := range sample{...} ***")
    for key, val := range sample {
        println(key, val)
    }
    println()

    println("*** p := sortedkeys(sample) ; for p.Range(){...} ***")
    p := sortedkeys.New(sample)
    for p.Range() {
        println(p.Key, p.Value)
    }
    println()

    println("*** for Go1.22 rangefunc ***")
    for key, val := range sortedkeys.Each(sample) {
        println(key, val)
    }
}
```

```go run example.go|
*** for key,val := range sample{...} ***
C 3
D 4
A 1
B 2

*** p := sortedkeys(sample) ; for p.Range(){...} ***
A 1
B 2
C 3
D 4

*** for Go1.22 rangefunc ***
A 1
B 2
C 3
D 4
```

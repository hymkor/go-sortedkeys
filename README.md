go-sortedkeys
=============

Make iterator for key-value pairs of map by sorted order

example
-------

```go
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

	println("*** p := sortedkeys(sample) ; for p.Range(){...} ***")
	p := sortedkeys.New(sample)
	for p.Range() {
		println(p.Key, p.Value)
	}
}
```

```
*** for key,val := range sample{...} ***
B 2
C 3
D 4
A 1
*** p := sortedkeys(sample) ; for p.Range(){...} ***
A 1
B 2
C 3
D 4
```

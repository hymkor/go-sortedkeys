// +build run

package main

import (
	"github.com/zetamatta/go-sortedkeys"
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

	println("*** for _,key := range sortedkeys.Strings(sample){...} ***")
	for _, key := range sortedkeys.Strings(sample) {
		println(key, sample[key])
	}
}

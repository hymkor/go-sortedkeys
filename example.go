//go:build run
// +build run

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

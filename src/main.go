package main

import (
	go_map "github.com/rusfort/go-sorted-map/map"
)

func main() {
	println("hello")

	gm := go_map.NewMap()
	m := gm.Map()
	m[123] = "hehe"
	m[7] = 4321
	m[-900] = -90.2

	println(gm.String())
	println(gm.Keys().String())
	println(gm.Size())

	m[12] = "abc123"

	println(gm.String())
	println(gm.Keys().String())
	println(gm.Size())
}


package main

import (
	"C"
	"eval/client"
	"unsafe"
)

//export test
func test(clientPtr uintptr) {
	client := (*client.Client)(unsafe.Pointer(clientPtr))
	client.Noop()
}

func main() {
}

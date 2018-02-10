package main

import (
	"fmt"

	ncs "github.com/hybridgroup/go-ncs"
)

func main() {
	res, name := ncs.GetDeviceName(0)
	if res != ncs.OK {
		fmt.Printf("NCS Error: %v\n", res)
	}

	fmt.Println("NCS: " + name)
}

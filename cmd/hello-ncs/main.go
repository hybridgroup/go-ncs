package main

import (
	"fmt"
	"os"
	"strconv"

	ncs "github.com/hybridgroup/go-ncs"
)

func main() {
	deviceID, _ := strconv.Atoi(os.Args[1])

	res, name := ncs.GetDeviceName(deviceID)
	if res != ncs.OK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	fmt.Println("NCS: " + name)

	// open device
	fmt.Println("Opening NCS device " + name + "...")
	d := ncs.OpenDevice(name)

	// close device
	fmt.Println("Closing NCS device " + name + "...")
	d.CloseDevice()

	fmt.Println("Done.")
}

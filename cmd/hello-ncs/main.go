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
	res, stick := ncs.OpenDevice(name)
	if res != ncs.OK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	// close device
	fmt.Println("Closing NCS device " + name + "...")
	stick.CloseDevice()

	fmt.Println("Done.")
}

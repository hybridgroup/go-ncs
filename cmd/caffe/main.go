package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	ncs "github.com/hybridgroup/go-ncs"
)

func main() {
	deviceID, _ := strconv.Atoi(os.Args[1])
	graphFileName := os.Args[2]

	res, name := ncs.GetDeviceName(deviceID)
	if res != ncs.OK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	fmt.Println("NCS: " + name)

	// open device
	fmt.Println("Opening NCS device " + name + "...")
	status, s := ncs.OpenDevice(name)
	if status != ncs.OK {
		fmt.Printf("NCS Error: %v\n", status)
		return
	}
	defer s.CloseDevice()

	// load graph file
	data, err := ioutil.ReadFile(graphFileName)
	if err != nil {
		fmt.Println("Error opening graph file:", err)
		return
	}

	// allocate graph
	fmt.Println("Allocating graph...")
	allocateStatus, graph := s.AllocateGraph(data)
	if allocateStatus != ncs.OK {
		fmt.Printf("NCS Error: %v\n", allocateStatus)
		return
	}

	// deallocate graph
	fmt.Println("Allocating graph...")
	graph.DeallocateGraph()

	fmt.Println("Done.")
}

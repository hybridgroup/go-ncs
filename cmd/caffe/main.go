package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"

	ncs "github.com/hybridgroup/go-ncs"
	"gocv.io/x/gocv"
)

func main() {
	deviceID, _ := strconv.Atoi(os.Args[1])
	graphFileName := os.Args[2]
	imageFileName := os.Args[3]

	res, name := ncs.GetDeviceName(deviceID)
	if res != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	fmt.Println("NCS: " + name)

	// open device
	fmt.Println("Opening NCS device " + name + "...")
	status, s := ncs.OpenDevice(name)
	if status != ncs.StatusOK {
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
	if allocateStatus != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", allocateStatus)
		return
	}

	// load image file
	img := gocv.IMRead(imageFileName, gocv.IMReadColor)

	// convert to format needed
	blob := gocv.BlobFromImage(img, 1.0, image.Pt(224, 224), gocv.NewScalar(104, 117, 123, 0), false, false)

	// load tensor into graph
	fmt.Println("Loading tensor...")
	fp16Blob := blob.ConvertFp16()
	loadStatus := graph.LoadTensor(fp16Blob.ToBytes())
	if loadStatus != ncs.StatusOK {
		fmt.Println("Error loading tensor data:", loadStatus)
		return
	}

	// deallocate graph
	fmt.Println("Deallocating graph...")
	graph.DeallocateGraph()

	fmt.Println("Done.")
}

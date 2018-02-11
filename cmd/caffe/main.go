// How to use:
//
// You must have OpenCV/GoCV installed to use this example.
// Run the following commands:
//
// 	source $GOPATH/src/gocv.io/x/gocv/env.sh
// 	go run ./cmd/caffe/main.go 0 ~/Development/ncsdk/examples/caffe/GoogLeNet/graph ~/Development/ncsdk/examples/data/images/cat.jpg ~/Development/ncsdk/examples/data/ilsvrc12/synset_words.txt
//
package main

import (
	"bufio"
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
	descriptions, _ := readDescriptions(os.Args[4])

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

	// convert to format needed by NCS
	resized := gocv.NewMat()
	gocv.Resize(img, resized, image.Pt(224, 224), 0, 0, gocv.InterpolationDefault)

	fp32Image := gocv.NewMat()
	resized.ConvertTo(fp32Image, gocv.MatTypeCV32F)

	fp16Blob := fp32Image.ConvertFp16()

	// load tensor into graph
	fmt.Println("Loading tensor...")
	loadStatus := graph.LoadTensor(fp16Blob.ToBytes())
	if loadStatus != ncs.StatusOK {
		fmt.Println("Error loading tensor data:", loadStatus)
		return
	}

	// get result
	resultStatus, data := graph.GetResult()
	if resultStatus != ncs.StatusOK {
		fmt.Println("Error getting results:", resultStatus)
		return
	}

	// convert results from fp16 back to float32
	fp16Results := gocv.NewMatFromBytes(1, len(data)/2, gocv.MatTypeCV16S, data)
	results := fp16Results.ConvertFp16()

	// determine the most probable classification
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(results)
	fmt.Printf("description: %v %v, maxVal: %v\n", maxLoc, descriptions[maxLoc.X], maxVal)

	// deallocate graph
	fmt.Println("Deallocating graph...")
	graph.DeallocateGraph()

	fmt.Println("Done.")
}

// readDescriptions reads the descriptions from a file
// and returns a slice of its lines.
func readDescriptions(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

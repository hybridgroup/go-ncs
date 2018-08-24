// How to use:
//
// This example opens a connection to a Movidius Neural Computer Stick (NCS)
// then uses OpenCV to open a camera, and start displaying the current classification
// of whatever the camera sees.
//
// You must have OpenCV/GoCV installed to use this example.
// Run the following commands:
//
// 	go run ./cmd/tf-video/main.go 0 0 ~/Development/ncsdk/examples/tensorflow/inception_v3/graph ~/Development/ncsdk/examples/tensorflow/inception_v3/categories.txt
//
package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"os"
	"strconv"

	ncs "github.com/hybridgroup/go-ncs"
	"gocv.io/x/gocv"
)

func main() {
	deviceID, _ := strconv.Atoi(os.Args[1])
	cameraID, _ := strconv.Atoi(os.Args[2])
	graphFileName := os.Args[3]
	descriptions, _ := readDescriptions(os.Args[4])

	// get name of NCS stick
	res, name := ncs.GetDeviceName(deviceID)
	if res != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	// open NCS device
	fmt.Println("Opening NCS device " + name + "...")
	status, s := ncs.OpenDevice(name)
	if status != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", status)
		return
	}
	defer s.CloseDevice()

	// load precompiled graph file in NCS format
	data, err := ioutil.ReadFile(graphFileName)
	if err != nil {
		fmt.Println("Error opening graph file:", err)
		return
	}

	// allocate graph on NCS stick
	fmt.Println("Allocating graph...")
	allocateStatus, graph := s.AllocateGraph(data)
	if allocateStatus != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", allocateStatus)
		return
	}
	defer graph.DeallocateGraph()

	webcam, err := gocv.VideoCaptureDevice(cameraID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", cameraID)
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Movidius Tensorflow Classifier")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	rgbImg := gocv.NewMat()
	defer rgbImg.Close()

	fp32Image := gocv.NewMat()
	defer fp32Image.Close()

	statusColor := color.RGBA{0, 255, 0, 0}

	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Error cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// convert image to half-float blob format needed by NCS
		fp16data := gocv.FP16BlobFromImage(img, 1.0/128.0, image.Pt(299, 299), 128.0, true, false)

		// load image tensor into graph on NCS stick
		loadStatus := graph.LoadTensor(fp16data)
		if loadStatus != ncs.StatusOK {
			fmt.Println("Error loading tensor data:", loadStatus)
			return
		}

		// get result from NCS stick in fp16 format
		resultStatus, data := graph.GetResult()
		if resultStatus != ncs.StatusOK {
			fmt.Println("Error getting results:", resultStatus)
			return
		}

		// convert results from fp16 back to float32
		fp16Results, _ := gocv.NewMatFromBytes(1, len(data)/2, gocv.MatTypeCV16S, data)
		results := fp16Results.ConvertFp16()

		// determine the most probable classification
		_, maxVal, _, maxLoc := gocv.MinMaxLoc(results)

		// display classification
		if maxLoc.X != -1 {
			desc := descriptions[maxLoc.X+1]
			info := fmt.Sprintf("description: %v, maxVal: %v", desc, maxVal)
			gocv.PutText(&img, info, image.Pt(10, img.Rows()/2), gocv.FontHersheyPlain, 1.2, statusColor, 2)
		}

		fp16Results.Close()
		results.Close()

		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
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

# go-ncs

## Go language bindings for the Movidius Neural Computer Stick

You must have a Movidius Neural Computer Stick (NCS) in order to use this package.

## Install

This package has only been tested on Ubuntu 16.04 LTS.

First, install the Movidius Neural Compute SDK from https://github.com/movidius/ncsdk

Once you have installed the SDK by following the instructions on the NCSDK repository using `make install` you can then download the compile the graph files for the Caffe GoogLeNet example by running the following commands:

    cd ./examples/caffe/GoogLeNet
    make prototxt
    make caffemodel
    make compile

This will download and compile the NCS graph file needed to run the examples using `go-ncs`.

## Using

Here is a very simple example of opening/closing a connection to an NCS stick:

```go
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
	if res != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	fmt.Println("NCS: " + name)

	// open device
	fmt.Println("Opening NCS device " + name + "...")
	res, stick := ncs.OpenDevice(name)
	if res != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	// close device
	fmt.Println("Closing NCS device " + name + "...")
	res = stick.CloseDevice()
	if res != ncs.StatusOK {
		fmt.Printf("NCS Error: %v\n", res)
		return
	}

	fmt.Println("Done.")
}
```

There are more examples in the `cmd` directory of this repository.

## Using go-ncs with GoCV

It is very useful to combine go-ncs with GoCV [(https://gocv.io)](https://gocv.io) to be able to use the video capture and processing abilities of GoCV along with the classification abilities of the NCS. Take a look at the `cmd\caffe` and `cmd\caffe-video` for examples.


## License

Licensed under the Apache 2.0 license. Copyright (c) 2018 The Hybrid Group.

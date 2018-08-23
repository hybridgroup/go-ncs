# go-ncs

This package contains Go language bindings for the Intel® Movidius™ Neural Compute Stick (NCS) [(https://developer.movidius.com/)](https://developer.movidius.com/).

You must have the Intel Movidius NCS hardware in order to use this package.

## Install

This package has been tested on Ubuntu 16.04 LTS and macOS "El Capitan" and "High Sierra".

You must first install the NCSDK before you can use the Intel Movidius Myriad 2 Neural Compute Stick. The official SDK only supports Linux. However, the fork created by [@milosgajdos83](https://github.com/milosgajdos83) has some initial support for macOS. Sorry, no Windows yet.

### macOS

The macOS support for the NCSDK currently is only for the API, not the graph compiler or other tools. To install, run the following commands:

    brew install coreutils opencv libusb pkg-config wget
    git clone https://github.com/milosgajdos83/ncsdk.git
    cd ncsdk
    git checkout macos-V1
    cd api/src && sudo make basicinstall

### Linux

You must have OpenCV and GoCV installed in order to use the Movidius SDK with Go:

https://gocv.io/getting-started/linux/

Once they are installed, you can run the following commands:

    git clone https://github.com/milosgajdos83/ncsdk.git
    cd ncsdk
    make install

## Compiling Models

If on Linux, once you have installed the SDK you can then download and compile the graph files for the Caffe GoogLeNet example by running the following commands:

    cd ./examples/caffe/GoogLeNet
    make prototxt
    make caffemodel
    make compile

This will download and compile the NCS graph file needed to run the `go-ncs` examples.

### Install the go-ncs Go package

Now you can install the go-ncs Go package:

    go get -d -u github.com/hybridgroup/go-ncs

Once you have installed `go-ncs` you can use it just like any other Golang package.

## Using

Here is a very simple example of opening/closing a connection to an NCS stick:

```go
package main

import (
	"fmt"

	ncs "github.com/hybridgroup/go-ncs"
)

func main() {
	res, name := ncs.GetDeviceName(0)
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
}
```

There are several examples in the `cmd` directory of this repository.

## Using go-ncs with GoCV

It is very useful to combine go-ncs with GoCV [(https://gocv.io)](https://gocv.io) to be able to use the video capture and processing abilities of GoCV along with the classification abilities of the NCS. Take a look at the `cmd\caffe` and `cmd\caffe-video` for examples.

## License

Licensed under the Apache 2.0 license. Copyright (c) 2018 The Hybrid Group.

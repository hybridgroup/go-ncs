package ncs

// #cgo LDFLAGS: -lmvnc
/*
#include <ncs.h>
*/
import "C"
import "unsafe"

// Status is the device status
type Status int

const (
	// StatusOK when the device is OK.
	StatusOK = 0

	// StatusBusy means device is busy, retry later.
	StatusBusy = -1

	// StatusError communicating with the device.
	StatusError = -2

	// StatusOutOfMemory means device out of memory.
	StatusOutOfMemory = -3

	// StatusDeviceNotFound means no device at the given index or name.
	StatusDeviceNotFound = -4

	// StatusInvalidParameters when at least one of the given parameters is wrong.
	StatusInvalidParameters = -5

	// StatusTimeout in the communication with the device.
	StatusTimeout = -6

	// StatusCmdNotFound means the file to boot Myriad was not found.
	StatusCmdNotFound = -7

	// StatusNoData means no data to return, call LoadTensor first.
	StatusNoData = -8

	// StatusGone means the graph or device has been closed during the operation.
	StatusGone = -9

	// StatusUnsupportedGraphFile means the graph file version is not supported.
	StatusUnsupportedGraphFile = -10

	// StatusMyriadError when an error has been reported by the device, use MVNC_DEBUG_INFO.
	StatusMyriadError = -11
)

// Stick
type Stick struct {
	DeviceHandle unsafe.Pointer
}

// Graph
type Graph struct {
	GraphHandle unsafe.Pointer
}

// GetDeviceName gets the name of the NCS stick located at index.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncGetDeviceName.html
//
func GetDeviceName(index int) (Status, string) {
	buf := make([]byte, 100)
	ret := Status(C.ncs_GetDeviceName(C.int(index), (*C.char)(unsafe.Pointer(&buf[0]))))
	return ret, string(buf)
}

// OpenDevice initializes an NCS device and returns a Stick.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncOpenDevice.html
//
func OpenDevice(name string) (Status, *Stick) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var deviceHandle unsafe.Pointer
	ret := C.ncs_OpenDevice(cName, &deviceHandle)
	return Status(ret), &Stick{DeviceHandle: deviceHandle}
}

// CloseDevice closes a previously opened NCS device.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncCloseDevice.html
//
func (s *Stick) CloseDevice() Status {
	res := C.ncs_CloseDevice(s.DeviceHandle)
	s.DeviceHandle = nil
	return Status(res)
}

// AllocateGraph allocates a graph for use on an NCS device, and
// returns a Graph.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncAllocateGraph.html
//
func (s *Stick) AllocateGraph(graphData []byte) (Status, *Graph) {
	var graphHandle unsafe.Pointer
	ret := Status(C.ncs_AllocateGraph(s.DeviceHandle, &graphHandle, unsafe.Pointer(&graphData[0]), C.uint(len(graphData))))
	return ret, &Graph{GraphHandle: graphHandle}
}

// DeallocateGraph deallocates and frees resources for a Graph.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncDeallocateGraph.html
//
func (g *Graph) DeallocateGraph() Status {
	return Status(C.ncs_DeallocateGraph(g.GraphHandle))
}

// LoadTensor starts inference on the NCS by providing input to the neural network.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncLoadTensor.html
//
func (g *Graph) LoadTensor(tensorData []byte) Status {
	return Status(C.ncs_LoadTensor(g.GraphHandle, unsafe.Pointer(&tensorData[0]), C.uint(len(tensorData))))
}

// GetResult retrieves the result of an inference that was previously initiated
// using the LoadTensor() method.
//
// For more information:
// https://movidius.github.io/ncsdk/c_api/mvncGetResult.html
//
func (g *Graph) GetResult() (Status, []byte) {
	resultData := C.struct_ResultData{}
	status := C.ncs_GetResult(g.GraphHandle, &resultData)
	data := C.GoBytes(unsafe.Pointer(resultData.data), C.int(resultData.length))
	return Status(status), data
}

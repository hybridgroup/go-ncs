package ncs

// #cgo LDFLAGS: -lmvnc
/*
#include <stdlib.h>
#include <core.h>
*/
import "C"
import "unsafe"

// Status is the device status
type Status int

const (
	// OK when the device is OK.
	OK = 0

	// Busy means device is busy, retry later.
	Busy = -1

	// Error communicating with the device.
	Error = -2

	// OutOfMemory means device out of memory.
	OutOfMemory = -3

	// DeviceNotFound means no device at the given index or name.
	DeviceNotFound = -4

	// InvalidParameters when at least one of the given parameters is wrong.
	InvalidParameters = -5

	// Timeout in the communication with the device.
	Timeout = -6

	// CmdNotFound means the file to boot Myriad was not found.
	CmdNotFound = -7

	// NoData means no data to return, call LoadTensor first.
	NoData = -8

	// Gone means the graph or device has been closed during the operation.
	Gone = -9

	// UnsupportedGraphFile means the graph file version is not supported.
	UnsupportedGraphFile = -10

	// MyriadError when an error has been reported by the device, use MVNC_DEBUG_INFO.
	MyriadError = -11
)

// GetDeviceName gets the name of the NCS stick at index.
func GetDeviceName(index int) (Status, string) {
	buf := make([]byte, 100)
	ret := Status(C.ncs_GetDeviceName(C.int(index), (*C.char)(unsafe.Pointer(&buf[0]))))
	return ret, string(buf)
}

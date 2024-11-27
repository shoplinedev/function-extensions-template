package function_extension_template

// #include <stdlib.h>
import "C"

import (
	"runtime"
	"unsafe"
)

// Log a message to the console using _Log.
func Log(message string) {
	ptr, size := StringToPtr(message)
	_Log(ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}

//export _Log
func _Log(ptr, size uint32)

func StringToPtr(s string) (uint32, uint32) {
	ptr := unsafe.Pointer(unsafe.StringData(s))
	return uint32(uintptr(ptr)), uint32(len(s))
}

package function_extensions_template

// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

func PtrToString(ptr uint32, size uint32) string {
	return unsafe.String((*byte)(unsafe.Pointer(uintptr(ptr))), size)
}

func StringToLeakedPtr(s string) (uint32, uint32) {
	size := C.ulong(len(s))
	ptr := unsafe.Pointer(C.malloc(size))
	copy(unsafe.Slice((*byte)(ptr), size), s)
	return uint32(uintptr(ptr)), uint32(size)
}

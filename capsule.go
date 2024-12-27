package python

/*
#define PY_SSIZE_T_CLEAN
#include <Python.h>



*/
import "C"
import (
	"reflect"
	"runtime"
	"unsafe"

	pointer "github.com/mattn/go-pointer"
)

var pinner runtime.Pinner

func NewCapsule(v interface{}, name *string, dtor unsafe.Pointer) (*Reference, error) {
	var nptr *C.char
	if name != nil {
		nptr = C.CString(*name)
		defer C.free(unsafe.Pointer(nptr))
	}
	// ptr := pointer.Save(v)
	pinner.Pin(v)
	ptr := reflect.ValueOf(v).UnsafePointer()
	if capsule := C.PyCapsule_New(ptr, nptr, nil); capsule != nil {
		return NewReference(capsule), nil
	} else {
		return nil, GetError()
	}
}

func (self *Reference) IsCapsule(name *string) bool {
	var nptr *C.char
	if name != nil {
		nptr = C.CString(*name)
		defer C.free(unsafe.Pointer(nptr))
	}
	return C.PyCapsule_IsValid(self.Object, nptr) != 0
}

func (self *Reference) GetPointer(name *string) unsafe.Pointer { // interface{} {
	if self.IsCapsule(name) {
		var nptr *C.char
		if name != nil {
			nptr = C.CString(*name)
			defer C.free(unsafe.Pointer(nptr))
		}
		ptr := C.PyCapsule_GetPointer(self.Object, nptr)
		return ptr
		// return pointer.Restore(ptr)
	} else {
		return nil
	}
}

func (self *Reference) Unref(name *string) {
	if self.IsCapsule(name) {
		var nptr *C.char
		if name != nil {
			nptr = C.CString(*name)
			defer C.free(unsafe.Pointer(nptr))
		}
		ptr := C.PyCapsule_GetPointer(self.Object, nptr)
		pointer.Unref(ptr)
	}
	self.Release()
	pinner.Unpin()
}

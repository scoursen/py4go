package python

/*
#define PY_SSIZE_T_CLEAN
#include <Python.h>



*/
import "C"
import "unsafe"

func NewCapsule(ptr unsafe.Pointer, name *string, dtor unsafe.Pointer) (*Reference, error) {
	var nptr *C.char
	if name != nil {
		nptr = C.CString(*name)
		defer C.free(unsafe.Pointer(nptr))
	}
	if capsule := C.PyCapsule_New(ptr, nptr, nil); capsule != nil {
		return NewReference(capsule), nil
	} else {
		return nil, GetError()
	}
}

func (ref *Reference) IsCapsule(name *string) bool {
	var nptr *C.char
	if name != nil {
		nptr = C.CString(*name)
		defer C.free(unsafe.Pointer(nptr))
	}
	return C.PyCapsule_IsValid(ref.Object, nptr) != 0
}

func (ref *Reference) GetPointer(name *string) unsafe.Pointer {
	if ref.IsCapsule(name) {
		var nptr *C.char
		if name != nil {
			nptr = C.CString(*name)
			defer C.free(unsafe.Pointer(nptr))
		}
		return C.PyCapsule_GetPointer(ref.Object, nptr)
	} else {
		return nil
	}
}

package python

/*
#define PY_SSIZE_T_CLEAN
#include <Python.h>
*/
import "C"

//
// Reference
//

type CPyObject = C.PyObject
type Reference struct {
	Object *CPyObject
}

func NewReference(pyObject *CPyObject) *Reference {
	return &Reference{pyObject}
}

func (self *Reference) Type() *Type {
	return NewType(self.Object.ob_type)
}

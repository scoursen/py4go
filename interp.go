package python

/*
#define PY_SSIZE_T_CLEAN
#include <Python.h>



*/
import "C"

type SubInterpreter struct {
	State *C.PyThreadState
}

func NewSubInterpreter() (*SubInterpreter, error) {
	return &SubInterpreter{C.Py_NewInterpreter()}, nil
}

func (self *SubInterpreter) End() {
	C.Py_EndInterpreter(self.State)
}

package api

// Here we add our "py_" C functions to the module

// Note that this file could have been combined with "py_.go", but we preferred a clean separation
// The cost is that we must forward-declare the "py_" functions in the cgo preamble here

import (
	python "github.com/scoursen/py4go"
)

/*
#cgo pkg-config: python3-embed

#define PY_SSIZE_T_CLEAN
#include <Python.h>

PyObject *py_api_sayGoodbye(PyObject *, PyObject *unused);
PyObject *py_api_concat(PyObject *, PyObject *args);
PyObject *py_api_concat_fast(PyObject *, PyObject **args, Py_ssize_t nargs);
*/
import "C"

func CreateModule() (*python.Reference, error) {
	if module, err := python.CreateModule("api"); err == nil {
		if err := module.AddModuleCFunctionNoArgs("say_goodbye", C.py_api_sayGoodbye); err != nil {
			module.Release()
			return nil, err
		}

		if err := module.AddModuleCFunctionArgs("concat", C.py_api_concat); err != nil {
			module.Release()
			return nil, err
		}

		if err := module.AddModuleCFunctionFastArgs("concat_fast", C.py_api_concat_fast); err != nil {
			module.Release()
			return nil, err
		}

		return module, nil
	} else {
		return nil, err
	}
}

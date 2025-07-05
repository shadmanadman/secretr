package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"fmt"
	"unsafe"

	"secretr/internal"
)

//export SecretrAdd
func SecretrAdd(name *C.char, secret *C.char, pass *C.char) *C.char {
	err := internal.StoreSecret(
		C.GoString(name),
		C.GoString(secret),
		C.GoString(pass),
	)
	if err != nil {
		return C.CString("error: " + err.Error())
	}
	return C.CString("ok")
}

//export SecretrGet
func SecretrGet(name *C.char, pass *C.char) *C.char {
	s, err := internal.RetrieveSecret(
		C.GoString(name),
		C.GoString(pass),
	)
	if err != nil {
		return C.CString("error: " + err.Error())
	}
	return C.CString(s)
}

//export SecretrList
func SecretrList() *C.char {
	names, err := internal.ListSecrets()
	if err != nil {
		return C.CString("error: " + err.Error())
	}
	var buf bytes.Buffer
	for _, name := range names {
		buf.WriteString(name + "\n")
	}
	return C.CString(buf.String())
}

//export SecretrDelete
func SecretrDelete(name *C.char) *C.char {
	err := internal.DeleteSecret(C.GoString(name))
	if err != nil {
		return C.CString("error: " + err.Error())
	}
	return C.CString("ok")
}

//export SecretrFree
func SecretrFree(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func main() {
	fmt.Println("OK")
}

//
// main.go
// Copyright (C) 2019 Tim Hughes
//
// Distributed under terms of the MIT license.
//
package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lperson
#include "person.h"
*/
import "C"
import (
	"fmt"
)

type (
	Person C.struct_APerson
)

func GetPerson(name string, long_name string) *Person {
	return (*Person)(C.get_person(C.CString(name), C.CString(long_name)))
}

func main() {
	var f *Person
	f = GetPerson("tim", "tim hughes")
	fmt.Printf("Hello Go world: My name is %s, %s.\n", C.GoString(f.name), C.GoString(f.long_name))
}

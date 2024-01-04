package main

import (
	//#cgo LDFLAGS : -ldl
	//#include "eval.h"
	//#include <stdlib.h>
	"C"
	"fmt"
	"os"
	"os/exec"
	"path"
	"unsafe"

	"github.com/google/uuid"

	"eval/client"
)

var TEMPLATE = `package main

import (
	"C"
	"eval/client"
	"unsafe"
)

//export eval
func eval(clientPtr uintptr) {
	client := (*client.Client)(unsafe.Pointer(clientPtr))
	%s
}

func main() {
}
`

func eval(client *client.Client, code string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	id := uuid.NewString()
	dir := path.Join(wd, id)
	if err := os.Mkdir(dir, 0700); err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	sourcePath := path.Join(dir, "main.go")
	execPath := path.Join(dir, "main\x00")
	if err := os.WriteFile(sourcePath, []byte(fmt.Sprintf(TEMPLATE, code)), 0666); err != nil {
		return err
	}
	cmd := exec.Command("go", "build", "-buildmode=c-shared", "-o", id, sourcePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	defer os.Remove(execPath)
	cpath := C.CString(execPath)
	defer C.free(unsafe.Pointer(cpath))
	C.eval(unsafe.Pointer(client), cpath)
	return nil
}

func main() {
	client := client.NewClient()

	code := `
    client.Noop()
    `

	if err := eval(client, code); err != nil {
		panic(err)
	}
}

package main

/*
#cgo LDFLAGS: -ldl

#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>


static void test(void* c) {
    void (*fn)();
    void *h = dlopen("eval.dylib", RTLD_LAZY);
    if (!h) {
        fprintf(stderr, "Error: %s\n", dlerror());
        return;
    }

    *(void**)(&fn) = dlsym(h, "test");
    if (!fn) {
        fprintf(stderr, "Error: %s\n", dlerror());
        dlclose(h);
        return;
    }

    fn(c);
    dlclose(h);
}
*/
import "C"
import (
	"eval/client"
	"unsafe"
)

func main() {
	c := client.Client{}
	C.test(unsafe.Pointer(&c))
}

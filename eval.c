#include "eval.h"

#include <dlfcn.h>
#include <stdio.h>

int eval(void* client, char* path) {
    void (*fn)(void* client);
    void* h = dlopen(path, RTLD_LAZY);
    if (!h) {
        fprintf(stderr, "error: %s\n", dlerror());
        return 1;
    }

    *(void**)(&fn) = dlsym(h, "eval");
    if (!fn) {
        fprintf(stderr, "error: %s\n", dlerror());
        dlclose(h);
        return 1;
    }

    fn(client);
    dlclose(h);
}

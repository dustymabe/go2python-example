#!/usr/bin/python
from cffi import FFI
ffibuilder = FFI()

ffibuilder.set_source("pyhello",
    """ //passed to the real C compiler
        #include "hello.h"
    """,
    extra_objects=["hello.so"])

ffibuilder.cdef("""
    extern void Hello();
    """)

if __name__ == "__main__":
    ffibuilder.compile(verbose=True)

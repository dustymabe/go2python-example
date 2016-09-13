
This directory contains the hello.go file as well as a vendored
version of github.com/fatih/color. We will generate a C shared library
by running:

```
$ go build -buildmode=c-shared -o hello.so .
$ ls
hello.so  hello.go  hello.h  README.md  vendor
```

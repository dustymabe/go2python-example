# go2python-example

```
sudo dnf install golang git python2 python2-cffi python2-devel redhat-rpm-config
export GOPATH=~/go
go get github.com/dustymabe/go2python-example
cd $GOPATH/src/github.com/dustymabe/go2python-example/
cd _Go/
go build -buildmode=c-shared -o hello.so .
cd ../_C/
gcc hello.c hello.so
LD_LIBRARY_PATH=$(pwd) ./a.out
cd ../_Python/
./hello_ffi_builder.py
LD_LIBRARY_PATH=$(pwd) ./hello.py
```


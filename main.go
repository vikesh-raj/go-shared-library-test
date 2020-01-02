package main

import (
    "C"
    "fmt"
    gopointer "github.com/mattn/go-pointer"
    "unsafe"
)

type Sample struct {
    name string
}

func NewSample(name string) Sample {
    return Sample{
        name: name,
    }
}

func (s *Sample) PrintSample() {
    fmt.Println("sample name = " + s.name)
}

//export NewSampleC
func NewSampleC(name *C.char) unsafe.Pointer {
    n := C.GoString(name)
    sample := NewSample(n)
    p := gopointer.Save(&sample)
    return p
}

//export PrintSampleC
func PrintSampleC(p unsafe.Pointer) {
    sample := gopointer.Restore(p).(*Sample)
    sample.PrintSample()
}

//export CloseSampleC
func CloseSampleC(p unsafe.Pointer) {
    gopointer.Unref(p)
}

func main() {}

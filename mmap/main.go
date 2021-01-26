package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const defaultMaxFileSize = 1 << 30
const defaultMemMapSize = 128 * (1 << 20)

type Demo struct {
	file    *os.File
	data    *[defaultMaxFileSize]byte
	dataRef []byte
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf(msg, v...))
	}
}

func (demo *Demo) mmap() {
	b, err := syscall.Mmap(int(demo.file.Fd()), 0, defaultMemMapSize, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	_assert(err == nil, "failed to mmap", err)
	demo.dataRef = b
	demo.data = (*[defaultMaxFileSize]byte)(unsafe.Pointer(&b[0]))
}

func (demo *Demo) grow(size int64) {
	if info, _ := demo.file.Stat(); info.Size() >= size {
		return
	}
	_assert(demo.file.Truncate(size) == nil, "Failed to truncate")
}

func (demo *Demo) munmap() {
	_assert(syscall.Munmap(demo.dataRef) == nil, "failed to mumap")
	demo.data = nil
	demo.dataRef = nil
}

func main() {

	_ = os.Remove("tmp.txt")
	f, _ := os.OpenFile("tmp.txt", os.O_CREATE|os.O_RDWR, 0644)
	demo := Demo{file: f}
	demo.grow(1)
	demo.mmap()
	defer demo.munmap()

	msg := "hello geektutu"
	demo.grow(int64(len(msg) * 2))

	for i, v := range msg {
		demo.data[i*2] = byte(v)
		demo.data[i*2+1] = byte(' ')
	}
	// at, _ := mmap.Open("./tmp.txt")
	// buff := make([]byte, 3)
	// _, _ = at.ReadAt(buff, 3)
	// _ = at.Close()
	// fmt.Println(string(buff))
}

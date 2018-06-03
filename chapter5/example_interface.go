package main

import (
	"fmt"
)

interface {
	Method01()
	Method02(i int) error
}

type Loader interface {
	Load(filename string) error
}

type ReaderWriter {
	io.Reader
	io.Writer
}

func main() {

}

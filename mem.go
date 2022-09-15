package main

import (
	"fmt"

	"github.com/c2fo/vfs/v6/vfssimple"
)

func main() {
	mem, err := vfssimple.NewLocation("mem:///tmp/")
	if err != nil {
		panic(err)
	}

	file, err := mem.NewFile("test.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	text := "testing"

	_, err = file.Write([]byte(text))
	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	data := make([]byte, len(text))
	_, err = file.Read(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

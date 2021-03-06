package Flist

import (
	"fmt"
	"os"
	"path/filepath"
)

func Glist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func Add(a, b int) int {
	c := a + b
	return c
}
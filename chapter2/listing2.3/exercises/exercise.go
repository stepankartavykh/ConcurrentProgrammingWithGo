package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func read_file(fileNamePath string) {
	dat, err := os.ReadFile(fileNamePath)
	if err != nil {
		panic(err)
	}
	fmt.Println(dat)
}

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(ex)
	fmt.Println(path)

	pwd, err := os.Getwd()
	fmt.Println(pwd)
}

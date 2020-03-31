package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Qs-F/walkup"
)

const target = "package.json"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ps := walkup.Walkup(wd, target, 0)
	if len(ps) < 1 {
		fmt.Println(".")
		os.Exit(1)
	}
	fmt.Println(filepath.Dir(ps[0]))
}

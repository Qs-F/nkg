package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Qs-F/walkup"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ps := walkup.Walkup(wd, "package.json", 0)
	if len(ps) < 1 {
		fmt.Println(".")
		os.Exit(1)
	}
	fmt.Println(ps[0])
}

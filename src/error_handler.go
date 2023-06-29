package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v \n", err)
		os.Exit(1)
	}
}

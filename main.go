package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// manage args
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <directory/archive>\n", os.Args[0])
		return
	}

	// figure out if its an archive or a directory.
	if strings.HasSuffix(os.Args[1], ".uzip") {
		err := unzip(os.Args[1])
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	err := zip(os.Args[1], os.Args[1]+".uzip")
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/AlperAkca79/mv/internal/mv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage of mv: \n\tmv oldFile newFile")
	} else {
		mv.Rename(os.Args[1], os.Args[2])
	}
}

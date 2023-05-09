package mv

import (
	"fmt"
	"os"
)

func Rename(oldFileName string, newFileName string) {

	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Print("File ", oldFileName, " has been named ", newFileName, " successfully.")
}

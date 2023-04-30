package writeToFile

import (
	"fmt"
	"log"
	"os"
)

func WriteToFile(content string) {
	fmt.Printf("Started\n")

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dirname = dirname + "/Dev/playgrounds/garbage"

	f, err := os.Create(dirname + "/writeToFile_test.txt")
	if err != nil {
		fmt.Printf("ERROR: %v", err)
	}
	f.WriteString(content)
	defer f.Close()

	fmt.Printf("All done!\n")
}

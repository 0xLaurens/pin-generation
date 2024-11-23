package main

import (
	"fmt"
	"log"
	"os"
)

const totalCombinations = 10_00_00_00 // 8 digits

func main() {
	filename := "pins.txt"
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}
	defer file.Close()

	for i := 0; i < totalCombinations; i++ {
		_, _ = file.WriteString(fmt.Sprintf("%08d\n", i))
	}
	log.Printf("Finished! saved combinations to %s", filename)
}

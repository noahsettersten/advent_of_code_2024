package utilities

import (
	"log"
	"os"
)

func ReadFile(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	converted := make([]string, len(input))
	line := ""

	for index, byte := range input {
		if byte == 10 {
			converted[index] = line
			line = ""
		} else {
			line += string(byte)
		}
	}

	return converted
}

// first := make([]string, 1000)
// second := make([]string, 1000)
// pair_index := 0
// current := ""
//
// for _, byte := range input {
// 	if byte == 10 {
// 		second[pair_index] = current
// 		pair_index++
// 		// fmt.Printf("Second of pair: %v\n", current)
// 		current = ""
// 		// println()
// 	} else if byte == 32 {
// 		if len(current) > 0 {
// 			first[pair_index] = current
// 			// fmt.Printf("First of pair: %v\n", current)
// 		}
// 		current = ""
// 		// print(" ")
// 	} else {
// 		current += string(byte)
// 		// print(byte)
// 	}
//
// }
//

package day1

import (
	"log"
	"os"
	"strconv"
)

func day1b() {
	input, err := os.ReadFile("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	first := make([]int, 1000)
	second := make([]int, 1000)
	pair_index := 0
	current := ""

	for _, byte := range input {
		if byte == 10 {
			val, err := strconv.Atoi(current)
			if err != nil {
				log.Fatal(err)
			}

			second[pair_index] = val
			pair_index++
			current = ""
		} else if byte == 32 {
			if len(current) > 0 {
				val, err := strconv.Atoi(current)
				if err != nil {
					log.Fatal(err)
				}

				first[pair_index] = val
			}
			current = ""
		} else {
			current += string(byte)
		}
	}

	similarity := 0

	for index := range first {
		left := first[index]
		count := 0

		for _, right := range second {
			if left == right {
				count++
			}
		}

		similarity += left * count
	}

	log.Printf("Similarity score: %v", similarity)
}

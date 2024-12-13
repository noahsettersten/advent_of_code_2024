package day1

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1a() {
	input, err := os.ReadFile("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	first := make([]string, 1000)
	second := make([]string, 1000)
	pair_index := 0
	current := ""

	for _, byte := range input {
		if byte == 10 {
			second[pair_index] = current
			pair_index++
			// fmt.Printf("Second of pair: %v\n", current)
			current = ""
			// println()
		} else if byte == 32 {
			if len(current) > 0 {
				first[pair_index] = current
				// fmt.Printf("First of pair: %v\n", current)
			}
			current = ""
			// print(" ")
		} else {
			current += string(byte)
			// print(byte)
		}

	}

	slices.SortFunc(first, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	slices.SortFunc(second, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})

	total := float64(0)

	for index := range first {
		fmt.Printf("Left %v, right %v\n", first[index], second[index])
		a, err := strconv.Atoi(first[index])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(second[index])
		if err != nil {
			log.Fatal(err)
		}

		distance := math.Abs(float64(a - b))
		println(distance)
		total += distance
		// distance = abs(distance)
		// println(value)
	}

	println(total)
	// 2192892

	// for _, value := range second {
	// 	println(value)
	// }

	// println(first)
	// println(second)
	// os.Stdout.Write(input)
	// print(input)
}

package day2

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/noahsettersten/advent_of_code_2024/utilities"
)

func processReport2(line string) bool {
	fields := strings.Split(line, " ")
	direction := 0
	previous := ""
	unsafe_levels := 0

	if len(line) == 0 {
		return false
	}

	print("\n ")
	print(line)
	print(" ")

	for _, field := range fields {
		if previous == "" {
			previous = field
		} else {
			prev, err := strconv.Atoi(previous)
			if err != nil {
				log.Fatal(err)
			}
			current, err := strconv.Atoi(field)
			if err != nil {
				log.Fatal(err)
			}

			if current > prev {
				print(" ", current, " > ", prev, " ")

				if direction == 0 {
					direction = 1
					print("++ ")
				} else if direction == 1 {
				} else {
					unsafe_levels++
				}
			} else if current < prev {
				print(" ", current, " < ", prev, " ")

				if direction == 0 {
					direction = -1
					print("-- ")
				} else if direction == -1 {
				} else {
					unsafe_levels++
				}
			} else {
				// No change.
				unsafe_levels++
			}

			difference := int(math.Abs(float64(prev) - float64(current)))

			if difference > 3 {
				// Too great of a difference
				unsafe_levels++
			} else {
				previous = field
			}

			if unsafe_levels > 1 {
				print("Not safe.")
				return false
			}
		}
	}

	print("Safe.")
	return true
}

func Day2b() {
	contents := utilities.ReadFile("./day2/day2_input.txt")
	safe_count := 0

	for _, line := range contents {
		safe := processReport2(line)

		if safe {
			safe_count++
		}
	}

	print("\n\n")
	log.Printf("Safe count: %v", safe_count)
}

package day2

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/noahsettersten/advent_of_code_2024/utilities"
)

func processReport(line string) bool {
	fields := strings.Split(line, " ")
	direction := 0
	previous := ""

	if len(line) == 0 {
		return false
	}

	print("\n ")
	print(line)
	print(" ")
	// log.Printf("%v", line)

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
				log.Printf("%v > %v", current, prev)

				if direction == 0 {
					direction = 1
					print("++ ")
				} else if direction == 1 {
				} else {
					// Direction changes, not safe
					print("Not safe. ->+")
					return false
				}
			} else if current < prev {
				log.Printf("%v < %v", current, prev)

				if direction == 0 {
					direction = -1
					print("-- ")
				} else if direction == -1 {
				} else {
					// Direction changes, not safe
					print("Not safe. +>-")
					return false
				}
			} else {
				// No change.
				print("Not safe. No change")
				return false
			}

			difference := int(math.Abs(float64(prev) - float64(current)))

			if difference > 3 {
				// Too great of a difference
				print("Not safe. Difference ")
				print(difference)
				return false
			} else {
				previous = field
			}
		}
	}

	print("Safe.")
	return true
}

func Day2a() {
	contents := utilities.ReadFile("./day2/day2_input.txt")
	safe_count := 0

	for _, line := range contents {
		safe := processReport(line)

		if safe {
			safe_count++
		}
	}

	print("\n\n")
	log.Printf("Safe count: %v", safe_count)
}

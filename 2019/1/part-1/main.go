package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	var (
		scanner = bufio.NewScanner(os.Stdin)
		fuel    = 0
	)
	for scanner.Scan() {
		v := scanner.Text()
		mass, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("invalid input: %v", err)
		}
		fuel += fuelRequired(mass)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(int(fuel))
}

func fuelRequired(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

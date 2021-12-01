package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func readInput(in io.Reader) ([]int, error) {
	var (
		nums []int
		r    = bufio.NewScanner(in)
	)
	for r.Scan() {
		t := strings.TrimSpace(r.Text())
		if t == "" {
			continue
		}
		n, err := strconv.Atoi(t)
		if err != nil {
			return nil, fmt.Errorf("failed to parse input: %w", err)
		}
		nums = append(nums, n)
	}
	if err := r.Err(); err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}
	return nums, nil
}

func run() error {
	const windowSize = 3
	var (
		last         int
		sum          int
		numIncreases int
	)
	nums, err := readInput(os.Stdin)
	if err != nil {
		return err
	}
	for i := range nums {
		if i > len(nums)-windowSize {
			break
		}
		sum = nums[i] + nums[i+1] + nums[i+2]

		// Do comparison only after we've filled the first window.
		if i > 0 && sum > last {
			numIncreases++
		}
		last = sum
	}
	fmt.Println(numIncreases)
	return nil
}

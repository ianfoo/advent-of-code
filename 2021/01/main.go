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
	var (
		last         int
		numIncreases int
	)
	nums, err := readInput(os.Stdin)
	if err != nil {
		return err
	}
	for i, num := range nums {
		if num > last && i > 0 {
			numIncreases++
		}
		last = num
	}
	fmt.Println(numIncreases)
	return nil
}

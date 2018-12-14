package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	freqs, err := list(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	repeated, iter := firstRepeatedSum(freqs)
	fmt.Printf("repeated %d on iteration %d\n", repeated, iter)
}

func list(r io.Reader) ([]int, error) {
	var (
		in    = bufio.NewScanner(r)
		freqs = make([]int, 0, 100)
	)
	for in.Scan() {
		ntxt := in.Text()
		n, err := strconv.Atoi(ntxt)
		if err != nil {
			return nil, err
		}
		freqs = append(freqs, n)
	}
	if err := in.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}

func firstRepeatedSum(list []int) (int, int) {
	var (
		sum, iter int
		hist      = make(map[int]struct{})
	)
	for {
		iter++
		for _, n := range list {
			sum += n
			if _, ok := hist[sum]; ok {
				return sum, iter
			}
			hist[sum] = struct{}{}
		}
	}
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("error: ")
}

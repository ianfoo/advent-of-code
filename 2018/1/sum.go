package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var sum int
	for in.Scan() {
		ntxt := in.Text()
		n, err := strconv.Atoi(ntxt)
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}
	if err := in.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

func init() {
	log.SetFlags(0)
	log.SetPrefix("error: ")
}

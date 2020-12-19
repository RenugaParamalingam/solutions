package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// a receives input. Use two channels named b and c.
	// b should print odd numbers. c should print even numbers.
	scanner := bufio.NewScanner(os.Stdin)

	b := make(chan int)
	c := make(chan int)

	go func() {
		for {
			select {
			case n := <-b:
				if n%2 != 0 {
					fmt.Println("odd ", n)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case n := <-c:
				if n%2 == 0 {
					fmt.Println("even ", n)
				}
			}
		}
	}()

	for scanner.Scan() {
		input := scanner.Text()

		a, err := strconv.Atoi(input)
		if err != nil {
			log.Println("error: ", err)
		}

		b <- a
		c <- a
	}

	close(b)
	close(c)
}

// Caitlin Schäffers
// Programmeertalen
// University of Amsterdam
//
// Hailstone sequence functions made with go routines.

package main

import (
	"os"
	"fmt"
	"strconv"
)

func hailstone_producer(n int, c chan int){
	// put all the calculates numbers in sequence in the channel
	for n > 1 {
		c <- n
		if n % 2 == 0 {
			n /= 2
		} else {
			n = n * 3 + 1
		}
	}
	// close the channel and signal that we're done when 1 is found
	close(c)
	fmt.Println("Done Producing!")
}

func hailstone_consumer(c chan int) {
	// make var to catch user input
    var dummy string

    // iterate over all values in channel
    for n := range c {
        fmt.Println("Press Enter to see the next value.")
        fmt.Scanln(&dummy)
        fmt.Printf("Received: %d\n", n)
    }
}

func main() {
	c := make(chan int)

	// if user hasn't given a start value for the sequence, exit progam
	if len(os.Args) != 2 {
		os.Exit(1)
		fmt.Println("Give a start number!")
	}

	// call the producer with the start value
	arg, _ := strconv.Atoi(os.Args[1])
	
	go hailstone_producer(arg, c)
	hailstone_consumer(c)
}

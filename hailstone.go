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

func hailstone_producer(n int) {
	counter := 1
	done := make(chan int)
	c := make(chan int)
	seq := make(chan int)
	var findNextHail func(n int)

	findNextHail = func(n int) {
		// put all the calculates numbers in sequence in the channel
		c <- n
		if n == 1 {
			// close the channel and signal that we're done when 1 is found
			close(c)
			done <- 1
		} else if n % 2 == 0 {
			seq <- (n / 2)
		} else {
			seq <- (n * 3 + 1)
		}
	}

	go findNextHail(n)

	// make anonymous function that catches the values put in the channels
	go func () {
		// count how many routines are still running
		for counter != 0 {
			select {

			case n := <-seq:
				counter++
				go findNextHail(n)

			case <- done:
				counter--
			}
		}
	}()

	hailstone_consumer(c)
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
	// if user hasn't given a start value for the sequence, exit progam
	if len(os.Args) != 2 {
		os.Exit(1)
		fmt.Println("Give a start number!")
	}

	// call the producer with the start value
	arg, err := strconv.Atoi(os.Args[1])
	if err == nil {
		hailstone_producer(arg)
	}
}


package main

import (
	"os"
	"fmt"
	"strconv"
)

 // PRODUCER MAKEN, die stuurt berekening naar hailstone consumer

  // MAIN die roept producer aan

func hailstone_producer(n int) {
	counter := 1
	done := make(chan int)
	c := make(chan int)
	seq := make(chan int)
	var findNextHail func(n int)

	findNextHail = func(n int) {
		c <- n
		if n == 1 {
			close(c)
			done <- 1
		} else if n % 2 == 0 {
			seq <- (n / 2)
		} else {
			seq <- (n * 3 + 1)
		}
	}

	go findNextHail(n)

	go func () {
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
    var dummy string

    for n := range c {
        fmt.Println("Press Enter to see the next value.")
        fmt.Scanln(&dummy)
        fmt.Printf("Received: %d\n", n)
    }
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
		fmt.Println("Give a start number!")
	}

	arg, err := strconv.Atoi(os.Args[1])
	if err == nil {
		hailstone_producer(arg)
	}
}

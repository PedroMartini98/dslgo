package main

/* Assignment
It's that time again, Mailio is hiring and we've been assigned to do the interview.
 The Fibonacci sequence is Mailio's interview problem of choice.
  We've been tasked with building a small toy program we can use in the interview.

Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci concurrently
Use a range loop to read from the channel and append the values to a slice
Return the slice */

func concurrentFib(n int) []int {
	// ?
	cf := make(chan int)
	go fibonacci(n, cf)

	output := []int{}
	for i := range cf {
		output = append(output, i)
	}
	return output
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

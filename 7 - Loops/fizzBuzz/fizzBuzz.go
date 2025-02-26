package main

import "fmt"

/* Assignment
We're hiring engineers at Textio, so time to brush up on the classic "Fizzbuzz" game,
a coding exercise that has been dramatically overused in coding interviews across the world.

Complete the fizzbuzz function that prints the numbers 1 to 100 inclusive each on their own line,
 but replace multiples of 3 with the text fizz and multiples of 5 with buzz. Print fizzbuzz for multiples of 3 AND 5.

Note: This lesson is graded based on the output of your program,
so don't leave any debugging print statements in your code. */

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("buzz")
		} else {

			fmt.Println(i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}

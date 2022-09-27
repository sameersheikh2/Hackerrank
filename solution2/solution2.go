package main

import "fmt"

func main() {
	var n, sum, num int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		sum += num
	}

	fmt.Print(sum)
}

// Okay, so the for loop looks little weird.  what := means. Basically it is declaring and initializing a variable n which is visible only in the scope of for loop. Next we have our loop conditions and increment operator. The parenthesis () around the whole thing is not required, however the {} are must even if you only have a single statement inside for loop. Then we are taking a integer input in num. Next We are adding that to sum.

// Hold on a minute, wouldn't sum += num throw a error for the first time? You never initialized the value of sum to 0. In golang it wouldn't. In golang all variables are assigned zero values when they are declared. Integers have 0 as zero values, strings have empty strings "" etc. So when we declared sum it also got initialized with 0, hence sum += num won't throw an error for the first time.

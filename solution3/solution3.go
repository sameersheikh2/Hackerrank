// Alice and Bob each created one problem for HackerRank. A reviewer rates the two challenges, awarding points on a scale from 1 to 100 for three categories: problem clarity, originality, and difficulty.

// The rating for Alice's challenge is the triplet a = (a[0], a[1], a[2]), and the rating for Bob's challenge is the triplet b = (b[0], b[1], b[2]).

// The task is to find their comparison points by comparing a[0] with b[0], a[1] with b[1], and a[2] with b[2].

// If a[i] > b[i], then Alice is awarded 1 point.
// If a[i] < b[i], then Bob is awarded 1 point.
// If a[i] = b[i], then neither person receives a point.
// Comparison points is the total points a person earned.

// Given a and b, determine their respective comparison points.
package main

import "fmt"

func main() {
	// declare 2 arrays of length 3 each and take inputs in those arrays.
	// This is the syntax for declaring an array in golang. Like variable, type comes later and we provide the length of array in []. Length of an array in golang is fixed.
	var a [3]int
	var b [3]int

	// referencing indexes 0,1,2 and taking inputs
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var ar, br int
	for i := 0; i < 3; i++ {
		//compare here
		if a[i] > b[i] {
			ar++
		} else if b[i] > a[i] {
			br++
		}
	}
	//  use of else if insted of else, need to take care of case where a[i]=b[i] no value should be incremented there
	fmt.Println(ar, br)
}

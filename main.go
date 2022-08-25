package main

import "fmt"

func main() {
	n := 7
	fmt.Println(numberOfMatches(n))

}
func numberOfMatches(n int) int {
	return roundCount(n, 0)
}

func roundCount(n, s int) int {
	sum := s
	temp := 0
	count := n / 2
	if n%2 == 0 {
		temp = count
	} else {
		temp = count + 1
	}
	sum = sum + count
	if n > 2 {
		sum = roundCount(temp, sum)
	}
	return sum
}

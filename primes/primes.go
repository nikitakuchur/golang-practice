package main

import "fmt"

func printPrimes(max int) {
outer:
	for n := 2; n <= max; n++ {
		if n == 2 {
			fmt.Println(n)
		}
		if n%2 == 0 {
			continue
		}
		for i := 3; i*i <= n; i += 2 {
			if n%i == 0 {
				continue outer
			}
		}
		fmt.Println(n)
	}
}

func main() {
	printPrimes(30)
}

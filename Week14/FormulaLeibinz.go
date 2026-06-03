package main

import "fmt"

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var pi float64
	for i := 1; i <= n; i++ {
		term := 1.0 / float64(2*i-1)
		if i%2 == 0 {
			term = -term
		}
		
		pi += term
	}
	pi *= 4

	fmt.Printf("Hasil PI: %.8f\n", pi)
}
package main

import "fmt"

func main() {
	var n int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	var piPrev, piCurr float64
	var selisih float64 = 1.0
	var i int

	for i = 1; i <= n && selisih > 0.00001; i++ {
		piPrev = piCurr 
		
		term := 1.0 / float64(2*i-1)
		if i%2 == 0 {
			term = -term
		}
		
		piCurr += 4 * term

		if i > 1 {
			selisih = piCurr - piPrev
			if selisih < 0 {
				selisih = -selisih
			}
		}
	}

	
	fmt.Printf("Hasil PI: %.10f\n", piPrev)
	fmt.Printf("Hasil PI: %.10f\n", piCurr)
	fmt.Printf("Pada i ke: %d\n", i)
}
package main

import (
	"fmt"
)

func GetMinMax(numbers ...*int) (min int, max int) {
	// your code here
	// Inisialisasi min dan max dengan nilai dari pointer pertama
	min = *numbers[0]
	max = *numbers[0]

	// Iterasi melalui array pointer untuk mencari nilai min dan max
	for _, num := range numbers {
		// Dereferencing pointer untuk mendapatkan nilai
		value := *num

		// Periksa nilai untuk menemukan nilai min dan max
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = GetMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Nilai max ", max)
}

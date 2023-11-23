package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	// Simpan nilai sementara dari *a ke dalam variable
	cup := *a

	// Ganti nilai *a dengan nilai *b
	*a = *b

	// Ganti nilai *b dengan nilai sementara (cup)
	*b = cup

	return *a, *b
}

func main() {
	a := 10
	b := 20

	Swap(&a, &b)
	fmt.Println(a, b)
}

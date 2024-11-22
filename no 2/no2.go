package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("A. Isi array:", array)

	fmt.Print("B. Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	fmt.Print("C. Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(array[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	fmt.Print("D. Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(array[i], " ")
		}
	}
	fmt.Println()

	var indexHapus int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indexHapus)

	if indexHapus >= 0 && indexHapus < n {
		array = append(array[:indexHapus],
			array[indexHapus+1:]...)
		n--
	}
	fmt.Println("E. Isi array setelah penghapusan:", array)

	var total int
	for _, value := range array {
		total += value
	}
	rataRata := float64(total) / float64(n)
	fmt.Printf("F. Rata-rata: %.2f\n", rataRata)

	var variance float64
	for _, value := range array {
		variance += math.Pow(float64(value)-rataRata, 2)
	}
	variance /= float64(n)
	standarDeviasi := math.Sqrt(variance)
	fmt.Printf("G. Standar deviasi: %.2f\n", standarDeviasi)

	var frekuensiBilangan int
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scan(&frekuensiBilangan)

	frekuensi := 0
	for _, value := range array {
		if value == frekuensiBilangan {
			frekuensi++
		}
	}
	fmt.Printf("H. Frekuensi dari bilangan %d: %d\n", frekuensiBilangan, frekuensi)
}

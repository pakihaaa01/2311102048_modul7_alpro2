package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	var pemenang []string
	for {
		var skorA, skorB int
		fmt.Print("skor pertandingan : ")
		fmt.Scan(&skorA, &skorB)
		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorA < skorB {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
	}

	fmt.Println("Hasil pertandingan:")
	for i, p := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, p)
	}
	fmt.Println("Pertandingan selesai")
}

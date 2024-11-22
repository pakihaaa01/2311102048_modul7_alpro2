package main

import "fmt"

type Titik struct {
	x, y int
}
type Lingkaran struct {
	tengah Titik
	jari   int
}

func kuadrat(n int) int {
	return n * n
}

func beradaDiDalamLingkaran(l Lingkaran, t Titik) bool {
	jarakKuadrat := kuadrat(t.x-l.tengah.x) + kuadrat(t.y-l.tengah.y)
	jariKuadrat := kuadrat(l.jari)
	return jarakKuadrat <= jariKuadrat
}
func main() {
	var l1, l2 Lingkaran
	var t Titik
	fmt.Scan(&l1.tengah.x, &l1.tengah.y, &l1.jari)
	fmt.Scan(&l2.tengah.x, &l2.tengah.y, &l2.jari)
	fmt.Scan(&t.x, &t.y)
	dalamL1 := beradaDiDalamLingkaran(l1, t)
	dalamL2 := beradaDiDalamLingkaran(l2, t)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

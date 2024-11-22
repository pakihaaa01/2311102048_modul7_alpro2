package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Masukan kata : ")
	fmt.Scanln(&input)
	*n = len(input)
	for i := 0; i < *n && i < NMAX; i++ {
		(*t)[i] = rune(input[i])
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Reverse: ")
	cetakArray(tab, m)
	if palindrom(tab, m) {
		fmt.Println("Palindrom: true")
	} else {
		fmt.Println("Palindrom: false")
	}
}

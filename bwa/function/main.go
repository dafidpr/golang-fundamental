package main

import "fmt"

func main() {
	sentence := printMyResult("Sayang sedang")
	fmt.Println(sentence)

	/* How to call a function that returns multiple returns, must create a variable according to the number of returns. How to call as below */
	luasPersegiPanjang, kelilingPersegiPanjang := calculateRectangle(10, 20)

	fmt.Println("Luas Persegi Panjang: ", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang: ", kelilingPersegiPanjang)

	luasPersegi, kelilingPersegi := calculateSquare(10)

	fmt.Println("Luas Persegi : ", luasPersegi)
	fmt.Println("Keliling Persegi : ", kelilingPersegi)
}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar Go"
	return newSentence
}

/* Multiple return (1)*/
func calculateRectangle(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

/* Multiple return (2)*/
func calculateSquare(sisi int) (luas int, keliling int) {
	luas = sisi * sisi
	keliling = 4 * sisi

	return
}

package main

import "fmt"

type LuasBangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {

	persegi := Persegi{Sisi: 4}
	luas := SeberapaLuas(persegi)

	fmt.Println("Luas Persegi : ", luas)

	persegiPanjang := PersegiPanjang{Panjang: 4, Lebar: 2}
	luas = SeberapaLuas(persegiPanjang)

	fmt.Println("Luas Panjang : ", luas)
}

func SeberapaLuas(bangunDatar LuasBangunDatar) int {
	return bangunDatar.HitungLuas()
}

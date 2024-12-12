package main

import (
	"fmt"
	"math"
)

type BangunDatar interface {
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

type Jajargenjang struct {
	alas   int
	tinggi int
}

func (jajargenjang Jajargenjang) HitungLuas() int {
	return jajargenjang.alas * jajargenjang.tinggi
}

type Segitiga struct {
	alas   int
	tinggi int
}

func (segitiga Segitiga) HitungLuas() int {
	return (segitiga.alas * segitiga.tinggi) / 2
}

type BelahKetupat struct {
	diagonal1 int
	diagonal2 int
}

func (belahKetupat BelahKetupat) HitungLuas() int {
	return belahKetupat.diagonal1 * belahKetupat.diagonal2 / 2
}

type Layang struct {
	diagonal1 int
	diagonal2 int
}

func (layang Layang) HitungLuas() int {
	return (layang.diagonal1 * layang.diagonal2) / 2
}

type Trapesium struct {
	a int
	b int
	t int
}

func (trapesium Trapesium) HitungLuas() int {
	return (trapesium.a + trapesium.b) * trapesium.t / 2
}

type Lingkaran struct {
	r float64
}

func (lingkaran Lingkaran) HitungLuas() float64 {
	return math.Pi * lingkaran.r * lingkaran.r
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

func main() {
	persegi := Persegi{Sisi: 5}
	luaspersegi := SeberapaLuas(persegi)
	fmt.Println("Luas Persegi : ", luaspersegi)

	persegiPanjang := PersegiPanjang{Panjang: 7, Lebar: 8}
	luaspersegipanjang := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luaspersegipanjang)

	segitiga := Segitiga{alas: 4, tinggi: 10}
	luassegitiga := SeberapaLuas(segitiga)
	fmt.Println("Luas Segitiga : ", luassegitiga)

	jajargenjang := Jajargenjang{alas: 9, tinggi: 5}
	luasjajargenjang := SeberapaLuas(jajargenjang)
	fmt.Println("Luas Jajargenjang : ", luasjajargenjang)

	belahketupat := BelahKetupat{diagonal1: 4, diagonal2: 4}
	luasbelahketupat := SeberapaLuas(belahketupat)
	fmt.Println("Luas Belah Ketupat : ", luasbelahketupat)

	lingkaran := Lingkaran{r: 1}
	luaslingkaran := lingkaran.HitungLuas()

	// Menampilkan hasil
	fmt.Printf("Luas Lingkaran dengan jari-jari %.2f adalah %.2f\n", lingkaran.r, luaslingkaran)

	trapesium := Trapesium{a: 5, b: 5, t: 6}
	luastrapesium := trapesium.HitungLuas()

	// menampilkan hasil
	fmt.Println("Luas Trapesium dengan jari-jari : ", luastrapesium)

}

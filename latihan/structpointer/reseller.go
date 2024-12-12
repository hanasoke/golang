package main

import "fmt"

type Reseller struct {
	Name     string
	Reseller []string
}

// AddGame
func (reseller *Reseller) AddProduct(product string) {
	reseller.Reseller = append(reseller.Reseller, product)
}

func main() {
	reseller := Reseller{Name: "Hanas"}

	reseller.AddProduct("Fast Cleansing Foam")
	reseller.AddProduct("Gel Brightening")
	reseller.AddProduct("Gel Acno")
	reseller.AddProduct("SC Serum")
	reseller.AddProduct("Brightening Serum")
	reseller.AddProduct("Soap 40 Gram")
	reseller.AddProduct("Soap 80 Gram")
	reseller.AddProduct("Soap Cepe 3")

	for _, product := range reseller.Reseller {
		fmt.Println(product)
	}

}

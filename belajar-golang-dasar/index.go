package main

import "fmt"

// func sayHelloTo(firstName string, lastName string) {
// 	fmt.Println("Hello", firstName, lastName)
// }

// func getHello(name string) string {
// 	if name == "" {
// 		return "Hello Bro"
// 	} else {
// 		return "Hello " + name
// 	}
// }

// func getFullName() (string, string, string) {
// 	return "Eko", "Kurniawan", "Khannedy"
// }

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"

	return
}

func main() {
	a, b, c := getFullName2()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// firstName, _, _ := getFullName()
	// fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
	// result := getHello("Eko")
	// fmt.Println(result)
	// fmt.Println(getHello(""))

	// firstName := "Eko"
	// lastName := "Kurniawan"
	// sayHelloTo(firstName, lastName)
	// sayHelloTo("Budi", "Nugraha")
	// for i := 0; i < 10; i++ {
	// 	sayHello()
	// }

	// sayHello()
	// sayHello()
	// sayHello()
	// fmt.Println("Eko")
	// fmt.Println("Eko Kurniawan")
	// fmt.Println("Eko Kurniawan Khannedy")

	// fmt.Println(len("Eko"))
	// fmt.Println("Eko Kurniawan"[4])
	// fmt.Println("Eko Kurniawan Khannedy"[7])

	// var name string

	// name = "Hanas Bayu Pratama"
	// fmt.Println(name)

	// name = "Saitama"
	// fmt.Println(name)

	// var friendName = "Budi"
	// fmt.Println(friendName)

	// var age = 30
	// fmt.Println(age)

	// country := "Indonesia"
	// fmt.Println(country)

	// var (
	// 	firstName = "Eko"
	// 	lastName  = "Khannedy"
	// )

	// fmt.Println(firstName)
	// fmt.Println(lastName)

	// const (
	// 	firstName string = "Eko"
	// 	lastName         = "Khannedy"
	// 	value            = 1000
	// )

	// fmt.Println(firstName)
	// fmt.Println(lastName)
	// fmt.Println(value)

	// var nilai32 int32 = 129
	// var nilai64 int64 = int64(nilai32)
	// var nilai8 int8 = int8(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai8)

	// var name = "Eko"
	// var e byte = name[0]
	// var eString string = string(e)

	// fmt.Println(name)
	// fmt.Println(eString)

	// var result = 10 - 13
	// fmt.Println(result)

	// var i = 10
	// i += 10 // i = i + 10
	// fmt.Println(i)

	// i++ // i = i + 1
	// fmt.Println(i)

	// var negative = -100
	// var positive = +100
	// fmt.Println(negative)
	// fmt.Println(positive)

	// var name1 = "Eko"
	// var name2 = "Saitama"
	// var result bool = name1 == name2
	// fmt.Println(result)

	// var value1 = 100
	// var value2 = 200

	// fmt.Println(value1 > value2)
	// fmt.Println(value1 < value2)
	// fmt.Println(value1 == value2)
	// fmt.Println(value1 != value2)

	// var names [3]string

	// names[0] = "Eko"
	// names[1] = "Kurniawan"
	// names[2] = "Kannedy"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[1])

	// var values = [3]int{
	// 	90,
	// 	95,
	// 	80,
	// }

	// fmt.Println(values)
	// fmt.Println(values[0])
	// fmt.Println(values[1])
	// fmt.Println(values[2])

	// fmt.Println(len(names))
	// fmt.Println(len(values))

	// var lagi [10]string

	// fmt.Println(len(lagi))

	// var months = [...]string{
	// 	"Januari",
	// 	"Februari",
	// 	"Maret",
	// 	"April",
	// 	"Mei",
	// 	"Juni",
	// 	"Juli",
	// 	"Agustus",
	// 	"September",
	// 	"Oktober",
	// 	"November",
	// 	"Desember",
	// }

	// var slice1 = months[4:7]
	// fmt.Println(slice1)
	// fmt.Println(len(slice1))
	// fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	// slice1[0] = "Mei Lagi"
	// fmt.Println(months)

	// var slice2 = months[11:]
	// fmt.Println(slice2)

	// var slice3 = append(slice2, "Eko")
	// fmt.Println(slice3)
	// slice3[1] = "Bukan Desember"
	// fmt.Println(slice3)

	// newSlice := make([]string, 2, 5)

	// newSlice[0] = "Eko"
	// newSlice[1] = "Kurniawan"

	// fmt.Println(newSlice)
	// fmt.Println(len(newSlice))
	// fmt.Println(cap(newSlice))

	// copySlice := make([]string, 1, cap(newSlice))
	// copy(copySlice, newSlice)
	// fmt.Println(copySlice)

	// iniArray := [5]int{1, 2, 3, 4, 5}
	// iniSlice := [...]int{1, 2, 3, 4, 5}

	// fmt.Println(iniArray)
	// fmt.Println(iniSlice)

	// runtime.GOMAXPROCS(2)

	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("Hanas")
	// go sayHelloTo("Samuel")
	// go sayHelloTo("Wirda")

	// var messages1 = <-messages
	// fmt.Println(messages1)

	// var messages2 = <-messages
	// fmt.Println(messages2)

	// var messages3 = <-messages
	// fmt.Println(messages3)

	// person := map[string]string{
	// 	"name":    "Eko",
	// 	"address": "Subang",
	// }

	// person["title"] = "Programmer"

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	// var book map[string]string = make(map[string]string)
	// book["title"] = "Belajar Go-Lang"
	// book["author"] = "Eko"
	// book["ups"] = "Salah"
	// fmt.Println(book)
	// fmt.Println(len(book))

	// delete(book, "ups")

	// fmt.Println(book)
	// fmt.Println(len(book))

	// var name = "Eko"

	// if name == "Eko" {
	// 	fmt.Println("Hello Eko")
	// } else if name == "Joko" {
	// 	fmt.Println("Hello Joko")
	// } else if name == "Budi" {
	// 	fmt.Println("Hello Budi")
	// } else {
	// 	fmt.Println("Hi, Kenalan Donk")
	// }

	// if length := len(name); length > 5 {
	// 	fmt.Println("Terlalu Panjang")
	// } else {
	// 	fmt.Println("Nama sudah benar")
	// }

	// name := "Kurniawan"

	// switch name {
	// case "Eko":
	// 	fmt.Println("Hello Eko")
	// 	fmt.Println("Hello Eko")
	// case "Joko":
	// 	fmt.Println("Hello Joko")
	// 	fmt.Println("Hello Joko")
	// default:
	// 	fmt.Println("Kenalan Donk")
	// 	fmt.Println("Kenalan Donk")
	// }

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// length := len(name)
	// switch {
	// case length > 10:
	// 	fmt.Println("Nama terlalu panjang")
	// case length > 5:
	// 	fmt.Println("Nama lumayan panjang")
	// default:
	// 	fmt.Println("Nama sudah benar")
	// }

	// for counter := 10; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	// slice := []string{"Eko", "Kurniawan", "Khannedy", "Budi", "Joko"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// for i, value := range slice {
	// 	fmt.Println("Index", i, "=", value)
	// }

	// person := make(map[string]string)
	// person["name"] = "Eko"
	// person["title"] = "Programmer"

	// for key, value := range person {
	// 	fmt.Println(key, "=", value)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println("Perulangan ke ", i)
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Perulangan ke", i)
	// }

}

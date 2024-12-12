package main

import "fmt"

func main() {
	// fmt.Println("Halo, belajar Golang")

	// add := calculation.Add(9, 9)
	// min := calculation.Min(18, 7)
	// kali := calculation.Kali(5, 7)
	// bagi := calculation.Bagi(10, 2)
	// even := calculation.Even(3)
	// odd := calculation.Odd(1)

	// fmt.Println(add)
	// fmt.Println(min)
	// fmt.Println(kali)
	// fmt.Println(bagi)
	// fmt.Println(even)
	// fmt.Println(odd)

	// var name string = "halo"
	// name = "Halo"
	// age := 20
	// age = 30
	// fmt.Println(name)
	// fmt.Println(age)

	// if
	// if else
	// else if
	// if bersarang

	// age := 9
	// if age > 10 {
	// 	fmt.Println("Boleh Bermain Game")
	// } else {
	// 	fmt.Println("Kamu belum boleh")
	// }

	// score := 67
	// var grade string
	// if score <= 50 {
	// 	grade = "E"
	// } else if score <= 60 {
	// 	grade = "D"
	// } else if score < 70 {
	// 	grade = "C"
	// } else {
	// 	grade = "NULL"
	// }
	// fmt.Println(grade)

	// value := 90
	// var grade string
	// if value <= 45 {
	// 	grade = "F"
	// } else if value <= 65 {
	// 	grade = "E"
	// } else if value <= 70 {
	// 	grade = "D"
	// } else if value <= 75 {
	// 	grade = "C"
	// } else if value <= 85 {
	// 	grade = "B"
	// } else if value <= 100 {
	// 	grade = "A"
	// } else {
	// 	grade = "NULL"
	// }
	// fmt.Println(grade)

	// number := 3
	// switch number {
	// case 1:
	// 	fmt.Println("Satu")
	// case 2:
	// 	fmt.Println("Dua")
	// case 3:
	// 	fmt.Println("Tiga")
	// default:
	// 	fmt.Println("DEFAULT")
	// }

	// aktifitas := 2
	// switch aktifitas {
	// case 1:
	// 	fmt.Println("Pagi")
	// case 2:
	// 	fmt.Println("Siang")
	// case 3:
	// 	fmt.Println("Sore")
	// case 4:
	// 	fmt.Println("Malam")
	// default:
	// 	fmt.Println("DEFAULT")
	// }

	// toko := 9
	// switch toko {
	// case 1:
	// 	fmt.Println("Toko Bunga")
	// case 2:
	// 	fmt.Println("Toko Lampu")
	// case 3:
	// 	fmt.Println("Toko Kelontong")
	// case 4:
	// 	fmt.Println("Toko Buah")
	// case 5:
	// 	fmt.Println("Toko Baja")
	// case 6:
	// 	fmt.Println("Warung Padang")
	// case 7:
	// 	fmt.Println("Warung Steak & Shake")
	// case 8:
	// 	fmt.Println("Toko SkinCare")
	// case 9:
	// 	fmt.Println("Tempat GYM")
	// }

	// fmt.Println("Saya sedang belajar Go")

	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar Go :", i)
	// 	i++
	// }

	// title := "Golang the best language"
	// for _, letter := range title {
	// 	fmt.Println("letter :", string(letter))
	// }
	// for index, letter := range title {
	// 	fmt.Println("index :", index, " letter :", string(letter))
	// }

	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index :", index, " letter :", string(letter))
	// 	}
	// }

	// for index, letter := range title {
	// 	letterString := string(letter)

	// 	if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
	// 		fmt.Println("index :", index, " letter :", string(letter))
	// 	}
	// 	switch letterString {
	// 	case "a", "i", "u", "e", "o":
	// 		fmt.Println("index :", index, " letter :", string(letter))
	// 	}
	// }

	// judul := "Hanas Bayu Pratama"
	// for _, letter := range judul {
	// 	fmt.Println("letter :", string(letter))
	// }
	// for index, letter := range judul {
	// 	fmt.Println("Index Ke -", index, "letter :", string(letter))
	// }

	// nama := "Hanas Bayu Pratama"
	// for index, letter := range nama {
	// 	letterString := string(letter)
	// 	// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
	// 	// 	fmt.Println("index :", index, "letter :", string(letter))
	// 	// }
	// 	// switch letterString {
	// 	// case "a", "i", "u", "e", "o":
	// 	// 	fmt.Println("index :", index, " letter :", string(letter))
	// 	// }
	// }

	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Pyhton"

	// languages := [...]string{
	// 	"Ruby",
	// 	"Pyhton",
	// 	"Javascript",
	// 	"Go",
	// 	"C#",
	// 	"Kotlin",
	// }

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)

	// for index, lang := range languages {
	// 	fmt.Println("Index : ", index, " language : ", lang)
	// }

	// days := [...]string{
	// 	"Senin",
	// 	"Selasa",
	// 	"Rabu",
	// 	"Kamis",
	// 	"Jum'at",
	// 	"Sabtu",
	// 	"Minggu",
	// }
	// fmt.Println(days)
	// length := len(days)
	// fmt.Println(length)
	// for index, lang := range days {
	// 	fmt.Println("Index : ", index, " Days : ", lang)
	// }

	// city := [...]string{
	// 	"Jakarta",
	// 	"Tanggerang",
	// 	"Surabaya",
	// 	"Denpasar",
	// 	"Bandung",
	// 	"Semarang",
	// 	"Yogyakarta",
	// 	"Balikpapan",
	// 	"Medan",
	// }
	// fmt.Println(city)
	// length := len(city)
	// fmt.Println(length)
	// for index, lang := range city {
	// 	fmt.Println("Index : ", index, " Days : ", lang)
	// }

	// var gamingConsoles []string
	// gamingConsoles = append(gamingConsoles, "PlayStation4")
	// gamingConsoles = append(gamingConsoles, "Nintendo Switch")
	// gamingConsoles = append(gamingConsoles, "Xbox One")
	// for _, console := range gamingConsoles {
	// 	fmt.Println(console)
	// }

	// var day []string
	// day = append(day, "Sunday")
	// day = append(day, "Monday")
	// day = append(day, "Tuesday")
	// day = append(day, "Wednesday")
	// day = append(day, "Thursday")
	// day = append(day, "Friday")
	// day = append(day, "Saturday")
	// for _, days := range day {
	// 	fmt.Println(days)
	// }

	// var myMap map[string]int
	// myMap = map[string]int{}
	// myMap["ruby"] = 9
	// myMap["Javascript"] = 8
	// myMap["go"] = 10
	// fmt.Println(myMap["Javascript"])

	// myMap := map[string]string{
	// 	"ruby":       "is beautiful",
	// 	"go":         "is super fast",
	// 	"javaScript": "hype",
	// }
	// value, isAvailable := myMap["ruby"]
	// fmt.Println(isAvailable)
	// fmt.Println(value)

	// for key, value := range myMap {
	// 	fmt.Println("Key : ", key, " value: ", value)
	// }
	// fmt.Println("==========")
	// delete(myMap, "ruby")
	// for key, value := range myMap {
	// 	fmt.Println("Key : ", key, " value: ", value)
	// }

	// days := map[string]string{
	// 	"sunday": "Holiday Day",
	// 	"monday": "Working Day",
	// }
	// fmt.Println(days)

	// students := []map[string]string{
	// 	{"name": "Agung", "score": "A"},
	// 	{"name": "Link", "score": "B"},
	// 	{"name": "Mario", "score": "E"},
	// }
	// // fmt.Println(students)\
	// for _, student := range students {
	// 	fmt.Println(student["name"], " - ", student["score"])
	// }

	// cars := []map[string]string{
	// 	{"brand": "Toyota", "name": "New Raize"},
	// 	{"brand": "Honda", "name": "Mobilio"},
	// 	{"brand": "Isuzu", "name": "D-Max"},
	// 	{"brand": "Mazda", "name": "2 Hatchback"},
	// 	{"brand": "Daihatsu", "name": "Xenia"},
	// }
	// fmt.Println(cars)
	// for _, car := range cars {
	// 	fmt.Println(car["brand"], " - ", car["name"])
	// }

	// // hitung rata-rata
	// scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// // Inisialisasi varibel total untuk menjumlahkan semua nilai
	// total := 0
	// // Menggunakan loop for untuk menjumlahkan semua nilai dalam array
	// for _, score := range scores {
	// 	total += score
	// }
	// // Menghitung rata - rata
	// average := float64(total) / float64(len(scores))
	// fmt.Printf("Rata - rata : %.2f\n", average)

	// scores := [9]int{100, 80, 75, 92, 70, 90, 88, 67, 96}
	// var goodScores []int
	// for _, score := range scores {
	// 	if score >= 90 {
	// 		goodScores = append(goodScores, score)
	// 	}
	// }
	// fmt.Println("Good Scores:", goodScores)

	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)
	// printMyResult("belajar")
	// printMyResult("Golang")
}

func printMyResult(sentence string) string {
	// fmt.Println(sentence)
	newSentence := sentence + " belajar Golang"
	return newSentence
}

// 1. input
// 2. proses
// 3. output

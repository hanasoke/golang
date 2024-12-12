package main

import (
	"fmt"
	"strconv"
)

// func sendData(ch chan<- int) {
// 	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
// 	for i := 0; true; i++ {
// 		ch <- i
// 		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
// 	}
// }
// func retrieveData(ch <-chan int) {
// loop:
// 	for {
// 		select {
// 		case data := <-ch:
// 			fmt.Print(`receive data "`, data, `"`, "\n")
// 		case <-time.After(time.Second * 5):
// 			fmt.Println("timeout. no activities under 5 seconds")
// 			break loop
// 		}
// 	}
// }
// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var messages = make(chan int)

// 	go sendData(messages)
// 	retrieveData(messages)
// }

// func getAverage(numbers []int, ch chan float64) {
// 	var sum = 0
// 	for _, e := range numbers {
// 		sum += e
// 	}
// 	ch <- float64(sum) / float64(len(numbers))
// }
// func getMax(numbers []int, ch chan int) {
// 	var max = numbers[0]
// 	for _, e := range numbers {
// 		if max < e {
// 			max = e
// 		}
// 	}
// 	ch <- max
// }
// func main() {
// 	runtime.GOMAXPROCS(2)
// 	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
// 	fmt.Println("numbers :", numbers)
// 	var ch1 = make(chan float64)
// 	go getAverage(numbers, ch1)
// 	var ch2 = make(chan int)
// 	go getMax(numbers, ch2)
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case avg := <-ch1:
// 			fmt.Printf("Avg \t: %.2f \n", avg)
// 		case max := <-ch2:
// 			fmt.Printf("Max \t: %d \n", max)
// 		}
// 	}
// }

// func sendMessage(ch chan<- string) {
// 	for i := 0; i < 20; i++ {
// 		ch <- fmt.Sprintf("data %d", i)
// 	}
// }
// func printMessage(ch <-chan string) {
// 	for message := range ch {
// 		fmt.Println(message)
// 	}
// }
// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var messages = make(chan string)
// 	go sendMessage(messages)
// 	printMessage(messages)
// }

// func main() {
// 	defer fmt.Println("halo")
// 	fmt.Println("selamat datang")
// }

// func main() {
// 	orderSomeFood("pizza")
// 	orderSomeFood("burger")
// }
// func orderSomeFood(menu string) {
// 	if menu == "pizza" {
// 		fmt.Print("Pilihan tepat! ", " ")
// 		fmt.Print("Burger ditempat kami paling enak!", "\n")
// 		return
// 	}

// 	fmt.Println("Pesanan anda : ", menu)
// }

func main() {
	// number := 3
	// if number == 3 {
	// 	fmt.Println("halo 1")
	// 	// defer fmt.Println("halo 3")
	// 	func() {
	// 		defer fmt.Println("halo 3")
	// 	}()
	// }
	// fmt.Println("halo 2")
	// os.Exit(1)
	// defer fmt.Println("halo")
	// fmt.Println("selamat datang")

	var input string
	fmt.Print("Type some number : ")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}

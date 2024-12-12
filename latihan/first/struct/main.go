package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) display() string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func displayUser(user User) string {
	return fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	// user := User{
	// 	Email:     "zelda@intendo.com",
	// 	FirstName: "Zeida",
	// 	LastName:  "Safitri",
	// 	ID:        1,
	// 	IsActive:  true}

	user := User{1, "Zeida", "Safitri", "zeldan@intendo.com", true}
	result := user.display()
	fmt.Println(result)

	user2 := User{2, "Link", "Awakening", "link@intendo.com", true}
	fmt.Println(user2.display())
	// users := []User{user, user2}

	// displayUser1 := displayUser(user)
	// displayUser2 := displayUser(user2)
	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)

	// group := Group{"Gamer", user, users, true}
	// displayGroup(group)
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

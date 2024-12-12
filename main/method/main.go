package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
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

func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func displayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member count : %d", len(group.Users))
	fmt.Println("")

	fmt.Println("Users name :")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func main() {
	user := User{1, "Zelda", "Safitri", "zelda@intendo.com", true}
	result := user.Display()

	fmt.Println(result)

	user2 := User{2, "Link", "Awakening", "link@intendo.com", true}
	fmt.Println(user2.Display())

	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}

	// displayGroup(group)
	group.DisplayGroup()
}

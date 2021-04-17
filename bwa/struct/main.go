package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

/* Method User */
func (user User) displayUser() string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

/* Method Group */
func (group Group) displayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Members count : %d", len(group.Users))
	fmt.Println("")

	for _, userGroup := range group.Users {
		fmt.Printf("Nama Member : %s %s", userGroup.FirstName, userGroup.LastName)
		fmt.Println("")
	}
}

func main() {
	/* Declaration User Struct */
	user := User{
		ID:        1,
		FirstName: "Dafid",
		LastName:  "Prasetyo",
		Email:     "dafid@gmail.com",
		IsActive:  true,
	}

	users := []User{user}
	group := Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}

	fmt.Println(user.displayUser())
	group.displayGroup()
}

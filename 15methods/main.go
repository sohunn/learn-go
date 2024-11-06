package main

import "fmt"

func main() {

	sohan := User{"sohan", "sohanshashi0884@gmail.com", true, 19}
	sohan.GetStatus()
	sohan.ChangeEmail("shamanth@gmail.com")
	fmt.Println(sohan.Email) // does not change, exactly why pointers come to use
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

func (u User) GetStatus() {
	fmt.Printf("Is %v verified? -> %v\n", u.Name, u.Verified)
}

func (u *User) ChangeEmail(newMail string) {
	u.Email = newMail
	fmt.Printf("Email successfully changed to %v\n", u.Email)
}

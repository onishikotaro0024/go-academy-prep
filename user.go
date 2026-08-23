package main

import "fmt"

type User struct {
	Name string
	Age  int
	Job  string
}

func printUser() {
	user := User{
		Name: "Kotaro",
		Age:  27,
		Job:  "Designer",
	}

	fmt.Println("名前:", user.Name)
	fmt.Println("年齢:", user.Age)
	fmt.Println("職業:", user.Job)
}

func (u User) Introduce() {
	fmt.Println("こんにちは、", u.Name, "です")
	fmt.Println("職業は:", u.Job, "です")
}

func (u User) GetJob() string {
	return u.Job
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) ChangeJob(newJob string) {
	u.Job = newJob
}

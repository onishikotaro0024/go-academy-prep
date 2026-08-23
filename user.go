package main

import "fmt"

type User struct {
	Name string
	Age  int
	Job  string
}

func (u User) Introduce() {
	fmt.Println("こんにちは、", u.Name, "です")
}

func (u User) GetJob() string {
	return u.Job
}

func (u *User) ChangeJob(newJob string) {
	u.Job = newJob
}

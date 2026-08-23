package main

import "fmt"

// type User struct {
// 	Name string
// 	Age  int
// 	Job  string
// }

// func (u User) Introduce() {
// 	fmt.Println("こんにちは、", u.Name, "です")
// }

// func (u User) GetJob() string {
// 	return u.Job
// }

// func (u *User) ChangeJob(newJob string) {
// 	u.Job = newJob
// }


type User struct {
	Name string
	Age  int
	Job  string
}


func name() {
	name := "Kotaro"
	age := 27

	fmt.Println(name)
	fmt.Println(age)

}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func checkAge(age int) {
	if age >= 65 {
		fmt.Println("シニアです")
	} else if age >= 20 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}
}


func printNames() {
	names := []string{"Kotaro", "ABC", "Taro"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	
names = append(names, "XYZ")
	
	for _, name := range names {
	fmt.Println(name)
}
}

func printProfile() {
	profile := map[string]string{
		"name": "Kotaro",
		"job":  "Engineer",
	}

	profile["country"] = "Japan"
	profile["job"] = "Designer"

	fmt.Println("名前:", profile["name"])
	fmt.Println("職業:", profile["job"]) // 「Designer」になる？
	fmt.Println("国:", profile["country"])


	email, ok := profile["email"]

if ok {
	fmt.Println(email)
} else {
	fmt.Println("emailは登録されていません")
}
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



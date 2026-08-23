package main

import "fmt"

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

func scopeTest() {
	name := "Kotaro"
	total := 0

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		total += i
	}

	fmt.Println(name)
	fmt.Println("合計:", total)
}

func printNames() {
	names := []string{"Kotaro", "ABC", "Taro"}

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
	fmt.Println("職業:", profile["job"])
	fmt.Println("国:", profile["country"])

	email, ok := profile["email"]

	if ok {
		fmt.Println(email)
	} else {
		fmt.Println("emailは登録されていません")
	}
}

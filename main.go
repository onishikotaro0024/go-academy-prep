package main

// パッケージ宣言

import "fmt"

// fmt パッケージ宣言

func main() {
	fmt.Println("Hello, Go!")

	name()

	greet("Kotaro")
	greet("ABC")

	result := add(10, 20)
	fmt.Println(result)

	checkAge(30)
	checkAge(18)

	repeatSum()

	scopeTest()

	printNames()

	printProfile()

	printUser() 
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

func add(a int, b int) int {
	return a + b
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

func repeatSum() {
	total := 0

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		total += i
	}

	fmt.Println("合計:", total)
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
	fmt.Println("職業:", profile["job"])　// 「Designer」になる？
	fmt.Println("国:", profile["country"])


	email, ok := profile["email"]

if ok {
	fmt.Println(email)
} else {
	fmt.Println("emailは登録されていません")
}
}

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

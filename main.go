package main

import (
	"fmt"

	"github.com/onishikotaro0024/go-academy-prep/calc"
)

// func main() {
// 	fmt.Println("Hello, Go!")

// 	user := User{
// 		Name: "Kotaro",
// 		Age:  27,
// 		Job:  "Designer",
// 	}

// 	user.Introduce()

// 	result := add(10, 20)
// 	fmt.Println(result)
// }



func main() {
	fmt.Println("Hello, Go!")

	name()

	greet("Kotaro")
	greet("ABC")

	resultAdd := calc.Add(10, 20)
	fmt.Println(resultAdd)


	resultDivide, err := calc.Divide(10, 2)

	if err != nil {
		fmt.Println("エラー: ", err)
	} else {
		fmt.Println("結果: ", resultDivide)
	}
	

	// checkAge(30)
	// checkAge(18)

	repeatSum()

	scopeTest()

	printNames()

	printProfile()

	printUser()

	user := User{
		Name: "Kotaro",
		Age:  27,
		Job:  "Designer",
	}
	user2 := User{
		Name: "Yamada",
		Age:  33,
		Job:  "Engineer",
	}
	user.Introduce()
	user2.Introduce()

	userJob := user.GetJob()
	fmt.Println("取得した職業:", userJob)

	userName := user.GetName()
	fmt.Println("取得した名前:", userName)

	fmt.Println("変更前:", user.Job)
	user.ChangeJob("Engineer")
	fmt.Println("変更後:", user.Job)

	pointer := &user
	fmt.Println(pointer.Name)
	fmt.Println(pointer.Job)
	user.ChangeJob("Director") //→ 少し遠回りだが、変更ルールをまとめられる
	//pointer.Job = "Director" → 速い・単純
	fmt.Println(user.Job)


	// result2, err2 := divide(10, 0)
	// if err2 != nil {
	// 	fmt.Println("エラー: ", err2)
	// } else {
	// 	fmt.Println("結果: ", result2)
	// }

	// err3 := checkNumber(-10)

	// if err3 != nil {
	// 	fmt.Println("エラー:", err3)
	// } else {
	// 	fmt.Println("正常です")
	// }

}

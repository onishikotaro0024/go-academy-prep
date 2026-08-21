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
	if age >= 65{
		fmt.Println("シニアです")
		}else if age >= 20{
		fmt.Println("成人です")
		}else{
		fmt.Println("未成年です")
		}
}

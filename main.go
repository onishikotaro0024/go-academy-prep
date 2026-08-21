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
}

func name() {
	name := "Kotaro"
	age := 30

	fmt.Println(name)
	fmt.Println(age)
}

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

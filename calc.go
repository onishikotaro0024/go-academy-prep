package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("0では割れません")
// 	}

// 	return a / b, nil
// }


func add(a int, b int) int {
	return a + b
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

func divide(a int, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("0では割れません")
	}
			return a / b, nil
}



func checkNumber(n int) error{
		if n < 0 {
			return fmt.Errorf("マイナスの値です")
	}
			return nil
}

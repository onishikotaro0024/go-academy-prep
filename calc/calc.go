package calc

import "fmt"

func Add(a int, b int) int {
	return a + b
}

// func Divide(a int, b int) (int, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("0では割れません")
// 	}

// 	return a / b, nil
// }

// func checkZero(b int) error {
// 	if b == 0 {
// 		return fmt.Errorf("0では割れません")
// 	}

// 	return nil
// }



func Divide(a int, b int) (int, error) {
		err := checkZero(b)
	if err != nil{
		return 0, err
		}
	return a / b, nil
}

func checkZero(b int) error {
	if b == 0{
		return fmt.Errorf("0では割れません")
		}
	return nil
}

// これでもOK↓//

// func Divide(a int, b int) (int, error) {
// 	checkZero := func(b int) error {
// 	if b == 0{
// 		return fmt.Errorf("0では割れません")
// 		}
// 	return nil
// 	}	
// 		err := checkZero(b)
// 	if err != nil{
// 		return 0, err
// 		}
// 	return a / b, nil
// }
	

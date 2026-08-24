package calc

func Add(a int, b int) int {
	return a + b
}


func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("0では割れません")
	}

	return a / b, nil
}

package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return a / b, nil
}

func main() {
	var a int = 3
	b := 3

	// go는 타입을 중요시 하기에 타입이 다를 경우 연산이 안됨
	fmt.Println(a + b)

	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)
	}

	fmt.Println(sum(a, b))

	result, err := divide(10, 0)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)

}

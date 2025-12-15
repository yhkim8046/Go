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

	// 타입과 길이가 고전인 배열
	var stricted [3]int = [3]int{1, 2, 3}

	fmt.Println(stricted)    // 1,2,3
	fmt.Println(stricted[0]) //1

	stricted[1] = 99

	fmt.Println(stricted[1])

	// 길이가 가변인 슬라이스
	s := []int{10, 20, 30}
	fmt.Println(s, len(s), cap(s))

	s = append(s, 40, 50)
	fmt.Println(s)

	// 자르기
	sub := s[1:4]
	fmt.Println(sub)

	//make 생성

	t := make([]int, 0, 5) // len = 0 cap =5
	t = append(t, 1, 2)
	fmt.Println(t, len(t), cap(t))
}

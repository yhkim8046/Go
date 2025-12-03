package main 

import "fmt"

func main(){
	var name string = "yuhwan"
	var age int = 27
    var role string = "Cloud Engineer"

	fmt.Println("Name: ", name)
	fmt.Println("age: ", age)
	fmt.Println("role: ", role)

	if role == "Cloud Engineer" {
		fmt.Println("True")
		return 
	}
	
	fmt.Println("Incorrect")
}
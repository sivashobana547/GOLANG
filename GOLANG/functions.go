package main

import "fmt"

func main() {
	greeting()

	name := "priya"
	namecall(name)

	add(10, 20)

	age := rfunc(0)
	fmt.Println("my age is", age)

	name1, age1 := bio("", 0)
	fmt.Printf("My name is %v and my age is %v\n", name1, age1)

	var num float64
	fmt.Print("Enter a num : ")
	fmt.Scan(&num)
	result := fact(num)
	fmt.Print("Enter a number to find its factorial: ", result)

}

func greeting() {
	fmt.Println("Welcome")
}

func namecall(name string) {
	fmt.Printf("Hello %v good morning\n", name)
}

func add(a int, b int) {
	var c int
	c = a + b
	fmt.Printf("The sum of %v and %v is %v\n", a, b, c)
}

func rfunc(age int) int {
	fmt.Print("enter your age : ")
	fmt.Scan(&age)
	return age
}

func bio(name1 string, age1 int) (string, int) {
	fmt.Print("enter your name : ")
	fmt.Scan(&name1)
	fmt.Print("enter your age : ")
	fmt.Scan(&age1)
	return name1, age1
}

//Recusion on Go
func fact(x float64) (y float64) {
	if x > 0 {
		y = x * fact(x-1)
	} else {
		y = 1
	}
	return
}

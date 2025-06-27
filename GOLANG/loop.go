package main
import "fmt"
func main(){
	names:=[]string{} .
	for { 
		var name string
		var age int

		fmt.Print("enter your name: ")
		fmt.Scan(&name)
		fmt.Print("enter your age: ")
		fmt.Scan(&age)
		fmt.Printf("I am %v and my age is %v\n", name, age)
		names=append(names,name)
		fmt.Println("Names : ", names)
	}
	
	for i:=0;i<=5;i++{ 
		fmt.Println(i)
	}

	
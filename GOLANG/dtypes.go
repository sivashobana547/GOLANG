package main
import "fmt"
func main(){
	var username string
	var age int
	username="hema"
	age=20
	fmt.Println(username)
	fmt.Println(age)
	fmt.Printf("Welcome %v and your age is %v",username,age)

	const n=10
	fmt.Printf("type of username is %T and type of age is %T and type of n is %T",username,age,n)
	fmt.Println()
	var num string = "547"
	fmt.Printf("type of num is %T",num)
	fmt.Println(
	var name string
	var age1 int
	fmt.Print("enter your name : ")
	fmt.Scan(&name)
	fmt.Print("enter your age : ")
	fmt.Scan(&age1)
	fmt.Printf("I am %v and my age is %v",name,age1)


}
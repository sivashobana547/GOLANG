package main
import "fmt"
func main(){
	var num int
	fmt.Print("Enter a number : ")
	fmt.Scan(&num)
	//if
	if num > 0 {
		fmt.Println("The number is positive")
	}
	//if...else
	if num%2==0{
		fmt.Println("The number is even")
	}else{
		fmt.Println("The unmber is odd")
	}
	//if...elseif...else
	if num <10{
		fmt.Println("The number is less than 5")
	}else if num< 20{
		fmt.Println("The number is less than 15")
	}else{
		fmt.Println("The number is greater than 25")
	}
}
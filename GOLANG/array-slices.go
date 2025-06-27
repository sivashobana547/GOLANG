package main

import "fmt"

func main() {
	syntax : var array_name [size(optional)]datatype = [size(optional)]datatype {value1, value2, ...}
	var book = []string{"IKIGAI", "The_Alchemist", "Atomic_Habits"}
	//array shuld contain same datatype
	fmt.Println(book)

	adding elements to an array
	/*var book1 string
	fmt.Println("Enter the name of the book:")
	fmt.Scan(&book1)
	book[3] = "The_power_of_now"
	book[4] = book1
	This will be used if the size of the array is fixed
	*/

	book = append(book, "I_am_That") // this will not work if the size is fixed
	fmt.Println(book)

	//slicing the array
	fmt.Printf("first book is %v \n", book[0])
	fmt.Printf("First three books are %v \n", book[0:3])
	fmt.Println("Length of array is", len(book))
	fmt.Println("Capacity of array is", cap(book))

	//Reverse a string
	s := "hello"
	lst := []rune(s)
	n := len(lst)
	for i := 0; i < n/2; i++ {
		lst[i], lst[n-1-i] = lst[n-1-i], lst[i]
	}
	fmt.Println("Reversed list is", string(lst))

	

	
}
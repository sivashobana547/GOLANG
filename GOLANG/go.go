package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Begins")
	go delay()
	fmt.Println("Ends")
    delaystr("Hello")
	go count(numch)
	a:=<-numch 
	fmt.Println(a)

	for a:=range numch{
		fmt.Println(a)
	}
}

func delay(){
	time.Sleep(2* time.Second)
	fmt.Println("Delay1 completed")
}

func delaystr(str string){
	for i:=0;i<5;i++{
		fmt.Println(str)
		time.Sleep(1* time.Second)
	}
	time.Sleep(2* time.Second)
	fmt.Println("Delay2 completed")
}

func count(ch chan int){
	for i:=0;i<5;i++{
		ch <- i 
	}
	
	close(ch) 
}
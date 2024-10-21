package main

import "fmt"
import "strconv"
func main(){
	for i:=1; i<=5; i++ {
		fmt.Println("This loop is " + strconv.Itoa(i) +" out of 5")
	}
	i := 1
	for i<5{
		fmt.Println("This loop is " + strconv.Itoa(i) +" out of 5")
		i++
	}
}

package main

import (
	"fmt"
)

func main() {

	//CONCATENATE STRINGS - By Ponni Sajeevan
	fmt.Println("\nCONCATENATE STRINGS - By Ponni Sajeevan")
	str1 := "'The course is WINP2000! "                                                        //String 1
	str2 := "The contributors are Ponni, Rose, Nikhitha, Dipanshu, Piyush, Swaroop, Abhishek'" //String 2
	fmt.Println("The concatanated string is:", Concat(str1, str2))                             //Print the output + Calling the function
}

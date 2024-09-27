package main

import (
	"fmt"
)

func main() {

	checkstr := "Welcome to Group A's Week 4 Project!"
	fmt.Println("Welcome to Group A's Week 4 Project!")

	//1. Concatanate Strings - Ponni Sajeevan
	fmt.Println("\nConcatanate Strings - By Ponni Sajeevan")
	str1 := "'The course is WINP2000! "                                                         //String 1
	str2 := "The contributors are Ponni, Rose, Nikhitha, Dipanshu, Piyush, Swaroop, Abhishek!'" //String 2
	fmt.Println("The concatenated string is:", Concat(str1, str2))                              //Print the output + Calling the function

	//2. Replace the characters in the string - Swaroop Krishna
	fmt.Println("\nReplace characters - By Swaroop")
	fmt.Println("Replacing the Characters from string:", ReplaceWord(checkstr, "t", "AA"))

	//3. Length of the string - Piyush
	fmt.Println("\nString Length - By Piyush")
	fmt.Println("The length of the string is:", StringLength(checkstr))

}

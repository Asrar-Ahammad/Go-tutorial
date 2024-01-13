package main

import "fmt"
import "unicode/utf8"
func main() {
	var intnums int = 324124
	fmt.Println(intnums)

	var floatNum float32
	floatNum = 2.34
	fmt.Println(floatNum)

	var myString string = "Helloř"
	fmt.Println(myString)
	fmt.Println(len(myString)) //? It is used to meaure the length of string. It Outputs the number of bytes
	fmt.Println(utf8.RuneCountInString("ř")) //? By using this we can get the exact length of a string

	var myRune rune = 'a'
	fmt.Println(myRune)

	var mybool bool = false
	fmt.Println(mybool)

	myvariable := "text"
	fmt.Println(myvariable)

	a,b,c := 1,2,3
	fmt.Println(a,b,c)

	const myconst string = "constant value"
	const pi float32 = 3.1415
	fmt.Println(pi)
}

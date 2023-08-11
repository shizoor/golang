//My notes so far on the Go language.
package main

import (
	"fmt"
	"strconv"
)

var message string = "hello there!"

// Main function
func main() {

	var i int = 1
	fmt.Println(message)
	for i <= 20 {
		fmt.Println(i)
		i++
	}
	var j int
	j = 42
	var k int = 27
	l := 99 //has to be a new variable, just use = to update
	fmt.Println(j + k + l)
	fmt.Printf("%v, %T, %v, %T, %v, %T\n", j, j, k, k, l, l)

	//declaring variables as a block
	var (
		color    string  = "red"
		radius   float32 = 6.7
		material         = "copper"
	)
	fmt.Println(color, radius, material)

	q := float64(radius)
	fmt.Println(q)

	// You can do a C style conversion of an int to a string, with the string acting as a char

	d := 83
	e := string(d)
	println(e)
	//  To do things nicer, as in to output numbers as strings, import strconv
	println("\n", strconv.Itoa(d))

	//booleans
	var flag bool = true
	var flag1 bool
	fmt.Printf("%v, %T", flag, flag)
	fmt.Printf("%v, %T", flag1, flag1)
	// == for comparisons as usual.

	//types : int8, int16, int32, int64, uint8, etc.
	//You need to do a type conversion between all different numerical types.

	//Bitwise operators

	ab := 8
	ac := 9
	fmt.Println(ab & ac)  // AND
	fmt.Println(ab & ac)  // OR
	fmt.Println(ab ^ ac)  // XOR ^ with no left operand for NOT  (powers are a function in the math package)
	fmt.Println(ab &^ ac) // AND NOT
	fmt.Println(ab >> 3)  // Bitshift right by 3
	fmt.Println(ab << 3)  // Bitshift left by 3  (doubles value each shift)

	//Scientific notation   No bitwise operators on non ints or non-bools.
	somenum := 3.14
	somenum = 13.7e72 //This overflows a float32
	somenum = 2.1e14
	fmt.Printf("%v, %T\n", somenum, somenum)

	//Go handles complex numbers!!!! also complex128 (2 float64s)
	var ni complex64 = 1 + 2i
	ni = complex(5, 6)
	fmt.Printf("%v, %T\n", ni, ni)
	fmt.Printf("%v, %T\n", real(ni), real(ni))
	fmt.Printf("%v, %T\n", imag(ni), imag(ni))

	//Characters in strings can be looked up.

	sometext := "I am a string, "
	fmt.Println(sometext)
	moretext := "I am another string."
	fmt.Println(sometext + moretext)

	fmt.Printf("%v, %T\n\n\n", string(sometext[7]), sometext[7]) //by default will output a uint8 without type conversion

	bytetext := []byte(moretext)
	fmt.Printf("%v, %T\n\n\n", bytetext, bytetext) //Convert to bytes.

	//To do utf-32 you need to use "runes".   Single quotes for a rune.

	myrune := '😁'
	fmt.Printf("%v, %T", myrune, myrune) //The smiley is returned as an int32 128513.

	//Use the string package and look into "ReadRune"

	//Constants

	const myConst int = 34 // remember all uppercase starting vars and consts get exported.
	// you can't use a function to declare a constant
	// Arrays are always variables.
	// iota increments each time used, and a block of const implies it so you don't have to type each time.

	const (
		firstconst = iota
		secondconst
		thirdconst
	)

	fmt.Println(firstconst, secondconst, thirdconst) //0 1 2

	const (
		KB = 1 << (10 * iota)
		MB
		GB
	)

	var filesize float32 = 4245765642
	fmt.Printf("%.2fGB", filesize/GB)

}

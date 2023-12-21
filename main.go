package main

import (
	"fmt"
	"time"
)

type Bob struct{
	name string
	age int
}

func main() {
	var numPtr *int
	numPtr = new(int) // returns pointer?
	fmt.Println(numPtr, *numPtr)

	// create bob and make him older
	bob := newBob("Benjamin", 24)
	fmt.Println(bob)
	bob.happyBirtday()
	fmt.Println(bob)

	// Just playing with imutability and links
	imstring := "Bob "
	fmt.Println(imstring, &imstring)

	imstring += "is happy "

	fmt.Println(imstring, &imstring)

	p := &imstring

	fmt.Println("Pointer: ", *p, p)

	*p += "again"

	fmt.Println("Pointer: ", *p, p, imstring, &imstring)	

	// Let's simulate Bob's life
	//lifeChannel := make(chan int)

	//go bob.timeFlow(lifeChannel)
	//getStatusAboutBobAge(lifeChannel)

}

func newBob(name string, age int) Bob {
	return Bob{name, age}
}

func (bob *Bob) happyBirtday() {
	bob.age++
}

func (bob *Bob) timeFlow(c chan int) {
	for {
		time.Sleep(1 * time.Second)
		bob.happyBirtday()
		c <- bob.age
	}	
}

func getStatusAboutBobAge(c chan int){
	for {
		age := <- c
		if age > 100{
			fmt.Printf("One more year's gone. But Bob is dead now...\n")
		} else {
			fmt.Printf("Congratulations! Bob is %v years old\n", age)
		}
	}
}




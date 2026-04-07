package main

import (
	"fmt"

	"github.com/demo/greetings"
)

func main() {

	str := greetings.Hello("Jimmy")

	fmt.Printf("Greeting message is : %s \n", str)

}

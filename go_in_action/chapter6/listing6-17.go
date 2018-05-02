package main
import (
	"fmt"
)

func main() {
	//unbuffered := make(chan int)
	buffered := make (chan string,10)
	buffered <- "Gopher"
	value := <- buffered
	fmt.Printf(value+"\n")
	fmt.Printf("Hello")
}

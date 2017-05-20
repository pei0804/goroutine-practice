package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	go process("A")
	go process("B")
	go process("C")
	fmt.Println("Finish!")
}

func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}

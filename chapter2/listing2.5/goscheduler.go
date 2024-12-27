package main

import (
	"fmt"
	"runtime"
)

/*
Note: this program is an example of what not to do; using go scheduler
to synchronize executions
*/
func sayHello() {
	fmt.Println("Inside")
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	fmt.Println("Mark 1 ====", runtime.NumGoroutine())
	go sayHello()
	fmt.Println("Mark 2 ====", runtime.NumGoroutine())
	// Question: What actually happenes here? Dive deep into!
	runtime.Gosched()
	fmt.Println("Mark 3 ====", runtime.NumGoroutine())
	fmt.Println("Finished")
}

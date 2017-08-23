package main

import "fmt"
import "time"

func longFonction(i int, done chan bool) {
	fmt.Printf("longFonction %d begins\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("longFonction %d ends\n", i)
	done <- true
}

func main() {
	fmt.Println("Bonjour.")

	maxBufSize, bufSize := 5, 0

	isDone := make(chan bool, maxBufSize)

	for i := 0; i < 20; i++ {
		go longFonction(i, isDone)
		bufSize++
		if bufSize >= maxBufSize {
			<-isDone
			bufSize--
		}
	}
	for bufSize > 0 {
		<-isDone
		bufSize--
	}

	/* 	fmt.Println("Press enter to quit...")
	   	var k string
	   	fmt.Scanln(&k)
	*/
}

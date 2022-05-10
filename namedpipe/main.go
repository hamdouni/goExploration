package main

import (
	"log"
	"os"
	"syscall"
	"time"
)

func main() {
	println("bonjour")

	var np = "/tmp/talker"

	syscall.Mkfifo(np, 0666)
	defer os.Remove(np)

	// pipe to write
	pipe, err := os.OpenFile(np, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}
	defer pipe.Close()

	//to open pipe to read
	file, err := os.OpenFile(np, os.O_RDONLY, os.ModeNamedPipe)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := []byte{' '}
	go func() {
		for {
			_, err := file.Read(buf)
			if err != nil {
				log.Printf("error while reading: %s", err)
			} else {
				log.Printf("read %c", buf)
			}
		}
	}()

	var val = []byte{'A'}
	for i := 0; i < 3; i++ {
		_, err := pipe.Write(val)
		if err != nil {
			log.Printf("error while writing %d: %s", i, err)
		}
		time.Sleep(1 * time.Second)
		val[0]++
	}

}

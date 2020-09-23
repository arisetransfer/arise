package main

import "fmt"
import "time"

func main() {
	//messages := make(chan string)
	done := make(chan bool)
	contents := make(map[string](chan []byte))
	contents["stuff"] = make(chan []byte)
	go func() {
		things["stuff"] <- []byte("1")
		time.Sleep(1 * time.Second)
		things["stuff"] <- []byte("2")
		time.Sleep(1 * time.Second)
		things["stuff"] <- []byte("3")
		time.Sleep(1 * time.Second)
		things["stuff"] <- []byte("4")
		time.Sleep(1 * time.Second)
		done <- true
	}()
L:
	for {
		select {
		case msg := <-things["stuff"]:
			fmt.Println(msg)
		case <-done:
			break L
		}
	}
}

/*package main

import {
	"fmt"
	"sync"
	"time"
}

const seats = 3
var process = new(sync.WaitGroup)

func client() {
    for {
        time.Sleep(3 * time.Second)
        clients <- &Client{}
    }
}


func barber(c chan chan string) {
	for {
		select{
		case visitor := <-c:
			name := <- visitor
			fmt.Print("%s is on ")
		}

	}


}





func main() {
	var clientNames []string = {"Andrew", "Daniel", "Ania", "Pasha", "Anton"}
	clientNumber = len(clientNames)
	process.Add(clientNumber)
	var WaitingRoom chan chan string
	go barber(WaitingRoom)
	for i := 0; i < clientNumber; i++ {
		time.Sleep(5 * time.Second)
		go (i)
	}
	process.Wait()
	fmt.Println("End of the day")
	time.Sleep(2 * time.Second)
}*/

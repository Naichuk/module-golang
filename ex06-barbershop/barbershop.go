package main

import (
	"fmt"
	"sync"
	"time"
)

const seats = 3

var process = new(sync.WaitGroup)

func client(Wr chan chan string, Cl chan string, name string) {
	flag := 0
	for {
		fmt.Printf("%s has come to the barbershop\n", name)
		select {
		case Wr <- Cl:
			Cl <- name
			<-Cl
			process.Done()
			flag = 1
		default:
			fmt.Printf("Waiting room is full, %s will come later\n", name)
			time.Sleep(5 * time.Second)
		}
		if flag == 1 {
			break
		}
	}
}

func barber(Wr chan chan string) {
	fmt.Printf("Welcome to Maksi's Awesome Barbershop!\n\n")
	for {
		select {
		case Cl := <-Wr:
			name := <-Cl
			fmt.Printf("%s is on haircutting now\n", name)
			time.Sleep(5 * time.Second)
			fmt.Printf("%ss haircut is done\n", name)
			Cl <- "ok"
		default:
			fmt.Printf("Barber Maksi is going to sleep\n")
			visitor := <-Wr
			fmt.Printf("Barber Maksi has woke up\n")
			name := <-visitor
			fmt.Printf("%s is on haircutting now\n", name)
			time.Sleep(8 * time.Second)
			fmt.Printf("%s haircut is done\n", name)
			visitor <- "done"
		}

	}

}

func main() {
	clientNames := [5]string{"Andrew", "Daniel", "Ania", "Pasha", "Anton"}
	clientNumber := len(clientNames)
	process.Add(clientNumber)
	WaitingRoom := make(chan chan string, seats)
	go barber(WaitingRoom)
	for i := 0; i < clientNumber; i++ {
		c := make(chan string)
		time.Sleep(3 * time.Second)
		go client(WaitingRoom, c, clientNames[i])
		time.Sleep(2 * time.Second)
	}
	process.Wait()
	fmt.Println("\n\nMaksi's Awesome Barbershop is closed. See you tomorrow!")
	time.Sleep(2 * time.Second)
}

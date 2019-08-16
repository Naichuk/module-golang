package main

import (
	"fmt"
	"sync"
	"time"
)

var waitg = new(sync.WaitGroup)

func Customer(lobby chan chan int, customer chan int, index int, names []string) {
	fmt.Println(names[index], " come to barbershop")
	select {
	case lobby <- customer:
		customer <- index
		<-customer
		waitg.Done()
	default:
		fmt.Println("Queue if full. ", names[index], "will come in 20 minutes")
		time.Sleep(10 * time.Second)
		Customer(lobby, customer, index, names)
	}
}

func Barbershop(lobby chan chan int, names []string) {
	for {
		select {
		case customer := <-lobby:
			index := <-customer
			name := names[index]
			fmt.Println(name, " has sat in barber's chair")
			time.Sleep(5 * time.Second)
			fmt.Println(name, "'s cut is done")
			customer <- 0
		default:
			fmt.Println("No one is in queue. Barber goes sleep")
			customer := <-lobby
			index := <-customer
			name := names[index]
			fmt.Println("Barber wake up")
			fmt.Println(name, " has sat in barber's chair")
			time.Sleep(7 * time.Second)
			fmt.Println(name, "'s cut is done")
			customer <- 0
		}
	}
}

func main() {
	customers := 10
	lobby := make(chan chan int, 2)
	names := []string{"Alex", "Ania", "Andrew", "Danya", "Maxim", "Artem", "Pasha", "Valera", "Tony", "Louis"}
	waitg.Add(customers)
	go Barbershop(lobby, names)
	for i := 0; i < customers; i++ {
		customer := make(chan int)
		time.Sleep(2 * time.Second)
		go Customer(lobby, customer, i, names)
	}
	waitg.Wait()
	fmt.Println("Barbershop closed. See ya !")
}

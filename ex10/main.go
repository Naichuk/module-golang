package main

import "time"

func main() {
	c := make(chan int)
	var l Node = NewNode()
	var f Node = NewNode()

	l.State = 1
	//l.AppendMessage("hello I'm leader", getSelfIp())
	l.AppendMessage("hello I'm leader")
	go Run(&l)
	f.SetMessage("Helloasasd")
	time.Sleep(4 * time.Second)
	go Send(&f)
	/*go StartListening()
	l.addVote()
	l.SetState(22)
	l.hasVoted()
	l.AppendMessage("hello")
	fmt.Println(l)
	var f Node = NewNode()
	f.SetState(123)
	f.Leader = &l
	fmt.Println(f.Leader.State)
	Send(&l)
	Send(&f)*/
	<-c
}

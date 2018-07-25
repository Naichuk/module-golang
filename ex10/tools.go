package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func getSelfIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	return addrs[1].String()
}

func Send(n *Node) {
	conn, err := net.Dial("tcp", "127.0.0.1:1337")
	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		tosend, _ := json.Marshal(n)
		s := string(tosend)
		conn.Write([]byte(s))
		conn.Close()
	}
}

func Listen(n *Node) {

	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		fmt.Println(err)
	}
	buf := make([]byte, 0, 256)
	tmp := make([]byte, 256)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//var n Node
		for {
			p, err := conn.Read(tmp)
			if err != nil {
				break
			}
			buf = append(buf, tmp[:p]...)
		}
		er := json.Unmarshal(buf, &n)
		fmt.Println(er)
		buf = buf[:0]
		//fmt.Println("Message Received:", n, "\n")
		//return n
		//if conn != nil {
		//go handleConnection(&conn)
	}
}

func handleConnection(c net.Conn) {
	var n Node
	//t := time.AfterFunc(MyTimeout*time.Second, printErr)
	buf := make([]byte, 256)
	c.Read(buf)
	//message, _ := ioutil.ReadAll(c)
	//ifmt.Println(message)
	//if message != "" {
	//if !t.Stop() {
	//	<-t.C
	//} else {
	//	t.Stop()
	//}
	//}
	//fmt.Println(time.Now())
	fmt.Println(buf)
	_ = json.Unmarshal(buf, &n)
	fmt.Println("Message Received:", n, "\n")
}

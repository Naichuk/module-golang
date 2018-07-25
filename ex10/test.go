package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

var mapIp = map[string]string{
	"10.30.8.105:1337": "Artem",
	"10.30.8.131:1337": "Yarik",
	"10.30.8.47:8888":  "Denis",
	"10.30.8.96:1337":  "Maks",
}

var Nodes []string

func main() {

	done := make(chan string)
	go StartListening()
	for key, _ := range mapIp {
		go Send(key)
	}
	<-done
}

func StartListening() {
	ln, err := net.Listen("tcp", ":1337")
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			// handle error
			fmt.Println(err)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message), "\n")
		//go handleConnection(conn)
	}
	/*tmp := make([]byte, 256)

	for {
		conn, _ := ln.Accept()
		for {
			n, err := conn.Read(tmp)
			if err != nil {
				break
			}
			buf = append(buf, tmp[:n]...)
		}
		var message Message
		err := json.Unmarshal(buf, &message)*/
}
func Send(str string) {
	for x := 0; x < 100; x++ {
		conn, err := net.Dial("tcp", str)
		time.Sleep(5 * time.Second)
		if err != nil {
			//fmt.Println(err)
			continue
		}
		tosend, _ := json.Marshal("hui" + mapIp[str])
		//tosend, _ := json.Marshal("")
		conn.Write(tosend)
		conn.Close()
		//fmt.Fprintf(conn, conn.LocalAddr().String())
	}

}

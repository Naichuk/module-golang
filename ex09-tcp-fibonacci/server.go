package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"time"
)

const tcpIp = "127.0.0.1:1234"

func fibonacci(num int64, x1 *big.Int, x2 *big.Int) *big.Int {
	if num == 0 {
		return x1
	}
	temp := big.NewInt(0)
	temp.Add(x1, x2)
	x1 = x2
	x2 = temp
	return fibonacci(num-1, x1, x2)
}

func Receive(conn net.Conn) {
	cash := make(map[int64]*big.Int)
	for {
		var msg int64
		var result string
		d := json.NewDecoder(conn)
		err := d.Decode(&msg)
		if err == io.EOF {
			conn.Close()
			return
		}
		fmt.Println("Received index:", msg)
		if num, ok := cash[msg]; !ok {
			x1 := big.NewInt(0)
			x2 := big.NewInt(1)
			start := time.Now()
			fbNum := fibonacci(msg, x1, x2)
			fmt.Println("Fibonacci number: ", fbNum)
			cash[msg] = fbNum
			time := time.Since(start)
			fmt.Println("Spended time: ", time)
			result = time.String() + " " + fbNum.String()
		} else {
			result = time.Duration(0).String() + " " + num.String()
		}
		Reply(conn, result)
		fmt.Printf("\n*waiting for new request...*\n")
	}
}

func Reply(conn net.Conn, answer string) {
	encoder := json.NewEncoder(conn)
	fmt.Println("Sending the response...")
	err := encoder.Encode(answer)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ln, err := net.Listen("tcp", tcpIp)
	if err != nil {
		fmt.Println("ERROR: Can't launch server")
		fmt.Println("ERROR: ", err)
	}
	fmt.Println("Launching server...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERROR: no data")
			fmt.Println("ERROR: ", err)
		}
		go Receive(conn)
	}
}

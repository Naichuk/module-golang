package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"time"
)

var IpAddr = "127.0.0.1:6666"

func Fibonacci(num int, x *big.Int, y *big.Int) *big.Int {
	if num == 0 {
		return x
	}
	temp := big.NewInt(0)
	temp.Add(x, y)
	x = y
	y = temp
	return Fibonacci(num-1, x, y)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Something wrong", err)
	}
}

func Receive(connection net.Conn, cash map[int]*big.Int) {
	var num int

	decoder := json.NewDecoder(connection)
	err := decoder.Decode(&num)
	if err != nil {
		CheckError(err)
		connection.Close()
		return
	}
	result := big.NewInt(0)
	result, ok := cash[num]
	var send string
	if ok != true {
		x := big.NewInt(0)
		y := big.NewInt(1)
		start := time.Now()
		result = Fibonacci(num, x, y)
		cash[num] = result
		end := time.Since(start)
		fmt.Println(end)
		send = end.String() + " " + result.String()
	} else {
		send = "0 s  " + result.String()
	}
	encoder := json.NewEncoder(connection)
	err = encoder.Encode(send)
	CheckError(err)
}
func main() {
	cash := make(map[int]*big.Int)
	listen, err := net.Listen("tcp", IpAddr)
	if err != nil {
		fmt.Println("Can't listen the port: ", err)
		return
	}
	for {
		fmt.Println("Listening port:6666")
		connection, err := listen.Accept()
		CheckError(err)
		Receive(connection, cash)
	}
}

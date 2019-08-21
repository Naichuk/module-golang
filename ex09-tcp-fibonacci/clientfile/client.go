package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

var IpAdd = "10.30.8.192:6666"

func CheckErr(err error) bool {
	if err != nil {
		fmt.Println("Something wrong: ", err)
		return true
	}
	return false
}

func main() {
	connection, err := net.Dial("tcp", IpAdd)
	if CheckErr(err) == true {
		return
	}
	defer connection.Close()
	fmt.Printf("Connected to 10.30.8.192\n")
	fmt.Printf("Enter the number ->")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if CheckErr(err) == true {
		return
	}
	num, err := strconv.Atoi(string(line))
	if CheckErr(err) == true {
		return
	}
	encoder := json.NewEncoder(connection)
	encoder.Encode(num)
	fmt.Println("Number sent")
	decoder := json.NewDecoder(connection)
	var fibb string
	err = decoder.Decode(&fibb)
	if CheckErr(err) == true {
		return
	}
	fmt.Println(fibb)
}

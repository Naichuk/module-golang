package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	const tcpIp = "127.0.0.1:1234"
	con, err := net.Dial("tcp", tcpIp)
	if err != nil {
		fmt.Println("ERROR: Connection failed")
		return
	}
	defer con.Close()
	for {
		var number string
		fmt.Println("Enter the index ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return
		}
		if input == "exit" {
			break
		}
		encodemsg := json.NewEncoder(con)
		sendmsg, err := strconv.ParseInt(input[:len(input)-1], 10, 32)

		if input == "exit" {
			break
		} else if err != nil {
			fmt.Println("ERROR: failed converting type string to integer(not a number)")
			break
		} else {
			encodemsg.Encode(sendmsg)
		}
		err = json.NewDecoder(con).Decode(&number)
		if err != nil {
			return
		}
		fmt.Println(number)
	}
}

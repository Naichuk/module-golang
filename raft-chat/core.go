package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Run(n *Node) {
	var currNode Node
	//fmt.Println("nullNode:", nullNode)
	go Listen(&currNode)
	//println(n.State)
	//fmt.Println(currNode.NewMessage)
	//fmt.Println(n)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$: ")
		scanner.Scan()
		text := scanner.Text()
		n.SetMessage(text)
		time.Sleep(1 * time.Second)
		//fmt.Println(currNode)
		//fmt.Println(n)
		switch currNode.State {
		case 0:
			if currNode.NewMessage != "" {
				n.Chain = append(n.Chain, currNode.NewMessage)
			}
			currNode.NewMessage = ""
			Send(n)
			currNode.State = 23
		case 1:
			//fmt.Println("I'm in case 1")
			if currNode.NewMessage != "" {
				currNode.AppendMessage(currNode.NewMessage)
			}
			n.UpdateChain(currNode.Chain)
			if n.NewMessage != "" {
				Send(n)
			}
			n.NewMessage = ""
			currNode.State = 23
			//currNode = nil
			/*case 0:
				if currNode.NewMessage != "" && n.State == 1 {
					n.Chain = append(n.Chain, currNode.NewMessage)
					fmt.Println(n.Chain)
					//time.Sleep(10 * time.Second)
					fmt.Println(currNode)
					currNode = nullNode
					fmt.Println(n)
				}
			case 1:
				currNode.UpdateChain(n.Chain)
				currNode = nullNode*/
		}
	}
}

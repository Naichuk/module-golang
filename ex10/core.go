package main

import "fmt"

func Run(n *Node) {
	var nullNode Node
	var currNode Node
	fmt.Println("nullNode:", nullNode)
	go Listen(&currNode)
	println(n.State)
	fmt.Println(currNode)
	if &currNode != nil {
		for {
			switch currNode.State {
			case 0:
				if currNode.NewMessage != "" && n.State == 1 {
					n.Chain = append(n.Chain, currNode.NewMessage)
					fmt.Println(n.Chain)
					currNode = nullNode
				}
			case 1:
				currNode.UpdateChain(n.Chain)
				currNode = nullNode
			}
		}
	}
}

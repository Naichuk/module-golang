package main

func Run(n *Node) {

	for {
		currNode := Listen()
		switch currNode.State {
		case 0:
			if currNode.NewMessage != "" && n.State == 1 {
				n.Chain = append(n.Chain, currNode.NewMessage)
			}
		case 1:
			currNode.UpdateChain(n.Chain)
		}
	}
}

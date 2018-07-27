package main

type Node struct {
	State      int
	Chain      []string
	NewMessage string
	LeaderAddr *Node
	Ip         string
	Votes      int
	Voted      bool
}
type LeadInt interface {
	AppendMessage(mes string)
	addVote()
}
type FollInt interface {
	UpdateChain(ch []string)
	SetMessage(mess string)
}

func (n *Node) UpdateChain(ch []string) {
	n.Chain = ch
}

func NewNode() Node {
	n := Node{0, make([]string, 0), "", nil, getSelfIp(), 0, false}
	return n
}

func (n *Node) SetState(st int) {
	n.State = st
}

func (n *Node) getIp() string {
	return n.Ip
}

func (n *Node) hasVoted() {
	n.Voted = true
}
func (n *Node) AppendMessage(mes string) {
	n.Chain = append(n.Chain, mes)
}
func (n *Node) addVote() {
	n.Votes++
}
func (n *Node) SetMessage(mess string) {
	n.NewMessage = mess
}

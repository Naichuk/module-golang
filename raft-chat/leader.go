package main

type Leader Node

type LeadInt interface {
	AppendMessage(mes string)
	addVote()
}

func (l *Leader) AppendMessage(mes string) {
	l.Chain = append(l.Chain, mes)
}
func (l *Leader) addVote() {
	l.Votes++
}

func (l *Leader) UpdateChain(ch []string) {
	l.Chain = ch
}

func (l *Leader) SetState(st int) {
	l.State = st
}

func (l *Leader) getIp() string {
	return l.Ip
}

func (l *Leader) hasVoted() {
	l.Voted = true
}

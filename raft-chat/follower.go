package main

type Follower Node

type FollInt interface {
	UpdateChain(ch []string)
	SetMessage(mess string)
}

func (f *Follower) UpdateChain(ch []string) {
	f.Chain = ch
}

func (f *Follower) SetState(st int) {
	f.State = st
}

func (f *Follower) getIp() string {
	return f.Ip
}

func (f *Follower) hasVoted() {
	f.Voted = true
}

func (f *Follower) SetMessage(mess string) {
	f.New_message = mess
}

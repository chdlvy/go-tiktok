package src

import "time"

type ModUser struct {
	id       uint64
	userName string
	password string
	Avatar   string
	Gender   uint8
	Birthday time.Time
}

func (this *ModUser) SetUserName(name string) {
	this.userName = name
}

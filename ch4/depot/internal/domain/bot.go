package domain

type BotStatus string

const (
	BotUnknown BotStatus = ""
	BotIdle    BotStatus = "idle"
	BotActive  BotStatus = "active"
)

func (s BotStatus) String() string {
	switch s {
	case BotIdle, BotActive:
		return string(s)
	default:
		return ""
	}
}

type Bot struct {
	ID     string
	Name   string
	Status BotStatus
}

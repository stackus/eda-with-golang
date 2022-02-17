package domain

type BotID string

type BotStatus string

const (
	BotUnknown BotStatus = ""
	BotIdle    BotStatus = "idle"
	BotActive  BotStatus = "active"
)

func (i BotID) String() string {
	return string(i)
}

func ToBotID(id string) BotID {
	return BotID(id)
}

func (s BotStatus) String() string {
	switch s {
	case BotIdle, BotActive:
		return string(s)
	default:
		return ""
	}
}

func ToBotStatus(status string) BotStatus {
	switch status {
	case BotIdle.String():
		return BotIdle
	case BotActive.String():
		return BotActive
	default:
		return BotUnknown
	}
}

type Bot struct {
	ID     string
	Name   string
	Status BotStatus
}

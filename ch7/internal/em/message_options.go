package em

type MessageOption interface {
	configureMessage(*message)
}

func (h Headers) configureMessage(msg *message) {
	msg.headers = h
}

type AckFn func() error

func (f AckFn) configureMessage(msg *message) {
	msg.ackFn = f
}

type NAckFn func() error

func (f NAckFn) configureMessage(msg *message) {
	msg.nackFn = f
}

type ExtendFn func() error

func (f ExtendFn) configureMessage(msg *message) {
	msg.extendFn = f
}

type KillFn func() error

func (f KillFn) configureMessage(msg *message) {
	msg.killFn = f
}

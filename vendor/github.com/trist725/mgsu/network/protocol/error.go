package protocol

import "fmt"

////////////////////////////////////////////////////////////////////////////////
type ErrUnsupportedMessage struct {
	MessageID interface{}
}

func (e ErrUnsupportedMessage) Error() string {
	return fmt.Sprintf("unsupported message, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrNoMessageHandler struct {
	MessageID interface{}
}

func (e ErrNoMessageHandler) Error() string {
	return fmt.Sprintf("no message handler, id=[%v]", e.MessageID)
}

////////////////////////////////////////////////////////////////////////////////
type ErrMessageTooLarge struct {
	MessageID interface{}
}

func (e ErrMessageTooLarge) Error() string {
	return fmt.Sprintf("message too large, id=[%v]", e.MessageID)
}

package internal

import (
	"mlgs/src/session"
)

type Logic struct {
	s *session.Session
}

func (logic *Logic) Init() error {
	logic.registerAllEventHandler()
	return nil
}

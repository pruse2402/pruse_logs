package services

import (
	"errors"

	"github.com/FenixAra/go-util/log"
)

type Ping struct {
	l *log.Logger
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func NewPing(l *log.Logger) *Ping {
	return &Ping{
		l: l,
	}
}

func (p *Ping) Ping() error {
	return nil
}

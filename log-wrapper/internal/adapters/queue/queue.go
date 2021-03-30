package queue

import (
	"github.com/FenixAra/go-util/log"
)

type Queuer interface {
	SendMessage(req *GelfReq) error
}

func New(l *log.Logger) (Queuer, error) {
	return &Kafka{
		l: l,
	}, nil
}

package queue

import (
	"pruse_logs/log-wrapper/dtos"
	"pruse_logs/log-wrapper/internal/adapters/queue"
	"pruse_logs/log-wrapper/internal/config/globals"

	"github.com/FenixAra/go-util/log"
)

type Queue struct {
	l     *log.Logger
	queue queue.Queuer
}

var reqChan = make(chan dtos.Log, 5000)

func New(l *log.Logger) (*Queue, error) {
	q, err := queue.New(l)
	if err != nil {
		return nil, err
	}
	return &Queue{
		l:     l,
		queue: q,
	}, nil
}

func (q *Queue) QueueLog(req *dtos.Log) error {
	reqChan <- *req
	return nil
}

func SendLogs() {
	q, err := queue.New(globals.Logger)
	if err != nil {
		globals.Logger.Errorf("Unable to get queue. Err: %+v", err)
	}

	for {
		req := <-reqChan
		gelfReq := queue.NewGelfReq(&req)
		q.SendMessage(gelfReq)
	}
}

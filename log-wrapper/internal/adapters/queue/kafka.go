package queue

import (
	"encoding/json"
	"fmt"
	"pruse_logs/log-wrapper/dtos"
	"pruse_logs/log-wrapper/internal/config"
	"pruse_logs/log-wrapper/internal/config/globals"
	"pruse_logs/log-wrapper/internal/services/prom"
	"strings"
	"time"

	"github.com/FenixAra/go-util/log"
	"github.com/Shopify/sarama"
)

type GelfReq struct {
	Line           string  `json:"_line,omitempty"`
	File           string  `json:"_file,omitempty"`
	RefID          string  `json:"_reference,omitempty"`
	FullMessage    string  `json:"full_message,omitempty"`
	Host           string  `json:"host"`
	Level          int     `json:"level"`
	ShortMessage   string  `json:"short_message"`
	Timestamp      float64 `json:"timestamp"`
	Version        string  `json:"version"`
	ResponseTime   float64 `json:"_response_time,omitempty"`
	StatusCode     int     `json:"_status_code,omitempty"`
	Method         string  `json:"_method,omitempty"`
	Request        string  `json:"_request,omitempty"`
	UserAgent      string  `json:"_user_agent,omitempty"`
	CustomerID     string  `json:"_customer_id,omitempty"`
	IPAddress      string  `json:"_ip_address,omitempty"`
	RequestGroup   string  `json:"_request_group,omitempty"`
	Severity       string  `json:"_severity,omitempty"`
	AppVersion     string  `json:"_app_version,omitempty"`
	TimeTaken      float64 `json:"_time_taken"`
	DependancyType string  `json:"_dependancy_type"`
	DependancyName string  `json:"_dependancy_name"`
}

func NewGelfReq(v *dtos.Log) *GelfReq {
	if v.Timestamp.IsZero() {
		v.Timestamp = time.Now()
	}

	if v.Title == "" && v.Message != "" {
		v.Title = v.Message
	}

	ut := float64(v.Timestamp.Unix())

	return &GelfReq{
		Version:        config.GELF_VERSION,
		Timestamp:      ut,
		ShortMessage:   v.Title,
		Level:          globals.LevelMap[strings.ToUpper(v.Level)],
		Host:           v.AppName,
		FullMessage:    v.Message,
		File:           v.File,
		Line:           v.Line,
		RefID:          v.RefID,
		ResponseTime:   v.ResponseTime,
		StatusCode:     v.StatusCode,
		Method:         v.Method,
		Request:        v.Request,
		UserAgent:      v.UserAgent,
		CustomerID:     v.CustomerID,
		IPAddress:      v.IPAddress,
		RequestGroup:   v.RequestGroup,
		Severity:       v.Level,
		AppVersion:     v.AppVersion,
		TimeTaken:      v.TimeTaken,
		DependancyType: v.DependancyType,
		DependancyName: v.DependancyName,
	}
}

var producer sarama.AsyncProducer

func init() {
	Init()
	go func() {
		for {
			select {
			case err := <-producer.Errors():
				globals.Logger.Errorf("Unable to send message to Producer. Err: %+v", err)
			}
		}
	}()
}

func Init() {
	var err error
	theArray := []string{"0.0.0.0:9092"}
	producer, err = sarama.NewAsyncProducer(theArray, nil)
	fmt.Println("", err)
	if err != nil {
		println("Unable to get new Async Producer. Err: %+v", err)
		return
	}
}

type Kafka struct {
	l *log.Logger
}

func (k *Kafka) SendMessage(req *GelfReq) error {
	t := time.Now()
	var status string
	var v time.Duration
	defer func() {
		go prom.RecordDependancyResponseTime("HTTP", "Kafka", status, v.Seconds())
	}()

	data, err := json.Marshal(req)
	if err != nil {
		k.l.Errorf("Unable to send message to Kafka. Err: %+v", err)
		v = time.Since(t)
		status = "Error"
		return err
	}

	// _, _, err = producer.SendMessage(&sarama.ProducerMessage{Topic: config.KAFKA_TOPIC, Key: nil, Value: sarama.ByteEncoder(data)})
	// if err != nil {
	// 	k.l.Errorf("Unable to send message to kafka queue. Err: %+v", err)
	// 	Init()
	// 	v = time.Since(t)
	// 	status = "Error"
	// 	return err
	// }

	producer.Input() <- &sarama.ProducerMessage{Topic: "testing", Key: nil, Value: sarama.ByteEncoder(data)}
	v = time.Since(t)
	status = "Success"
	return nil
}

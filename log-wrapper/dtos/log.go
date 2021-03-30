package dtos

import "time"

type Log struct {
	Level          string    `json:"level" example:"INFO" binding:"required"`
	Timestamp      time.Time `json:"timestamp" example:"2018-08-09T15:46:29.000Z" binding:"required"`
	Title          string    `json:"message" example:"message" binding:"required"`
	Message        string    `json:"full_message" example:"full message with details" binding:"required"`
	AppName        string    `json:"app_name" example:"Booking" binding:"required"`
	RefID          string    `json:"ref_id" example:"RefID" binding:"required"`
	File           string    `json:"file" example:"app file name" binding:"required"`
	Line           string    `json:"line" example:"line of file" binding:"required"`
	ResponseTime   float64   `json:"response_time" example:"1.012"`
	StatusCode     int       `json:"status_code" example:"200"`
	Method         string    `json:"method" example:"POST/GET"`
	Request        string    `json:"request" example:"/ping"`
	UserAgent      string    `json:"user_agent" example:"ios/android"`
	CustomerID     string    `json:"customer_id" example:"CustomerID"`
	IPAddress      string    `json:"ip_address" example:"127.0.0.1"`
	RequestGroup   string    `json:"request_group" example:"Ping"`
	AppVersion     string    `json:"app_version" example:"App Version"`
	TimeTaken      float64   `json:"time_taken" example:"1.11"`
	DependancyType string    `json:"dependancy_type" example:"http,database"`
	DependancyName string    `json:"dependancy_name" example:"googleapi,booktripsp"`
}

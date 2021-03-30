package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FenixAra/go-util/log"
)

const (
	ERR_MSG = "ERROR_MESSAGE"
	MSG     = "MESSAGE"
)

type RequestData struct {
	l     *log.Logger
	Start time.Time
	w     http.ResponseWriter
	r     *http.Request
}

type RenderData struct {
	Data  interface{}
	Paths []string
}

type TemplateData struct {
	Data interface{}
}

func (t *TemplateData) SetConstants() {

}

func logAndGetContext(w http.ResponseWriter, r *http.Request) *RequestData {
	w.Header().Add("X-Content-Type-Options", "nosniff")
	w.Header().Add("X-Frame-Options", "DENY")
	ref := r.Header.Get("ReferenceID")

	cfg := log.NewConfig(ref, "Debug", "Full", "config.APP_NAME", "", "", "")

	l := log.New(cfg)
	l.Info("Serving Request:", r.RequestURI, ", Method: ", r.Method)
	start := time.Now()

	return &RequestData{
		l:     l,
		Start: start,
		r:     r,
		w:     w,
	}
}

func redirectTo(path string, rd *RequestData) {
	rd.l.Info("Status Code:", http.StatusFound, ", Response time:",
		time.Since(rd.Start), ", Response: url redirect - ", path)
	http.Redirect(rd.w, rd.r, path, http.StatusFound)
}

func jsonifyMessage(msg string, msgType string, httpCode int) ([]byte, int) {
	var data []byte
	var Obj struct {
		Status   string `json:"status"`
		HTTPCode int    `json:"httpCode"`
		Message  string `json:"message"`
	}
	Obj.Message = msg
	Obj.HTTPCode = httpCode
	switch msgType {
	case ERR_MSG:
		Obj.Status = "FAILED"

	case MSG:
		Obj.Status = "SUCCESS"
	}
	data, _ = json.Marshal(Obj)
	return data, httpCode
}

func writeJSONMessage(msg string, msgType string, httpCode int, rd *RequestData) {
	d, code := jsonifyMessage(msg, msgType, httpCode)
	writeJSONResponse(d, code, rd)
}

func writeJSONStruct(v interface{}, code int, rd *RequestData) {
	d, err := json.Marshal(v)
	if err != nil {
		writeJSONMessage("Unable to marshal data. Err: "+err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		return
	}
	writeJSONResponse(d, code, rd)
}

func writeJSONResponse(d []byte, code int, rd *RequestData) {
	if code == http.StatusInternalServerError {
		rd.l.Info("Status Code:", code, ", Response time:", time.Since(rd.Start), " Response:", string(d))
	} else {
		rd.l.Info("Status Code:", code, ", Response time:", time.Since(rd.Start))
	}

	rd.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	rd.w.Header().Set("Access-Control-Allow-Origin", "*")
	rd.w.WriteHeader(code)
	rd.w.Write(d)
}

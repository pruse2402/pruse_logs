package handlers

import (
	"encoding/json"
	"net/http"

	"pruse_logs/log-wrapper/dtos"
	"pruse_logs/log-wrapper/internal/services/prom"
	"pruse_logs/log-wrapper/internal/services/queue"

	"github.com/julienschmidt/httprouter"
)

func setLogRoutes(router *httprouter.Router) {
	router.POST("/v1/logs", QueueLog)
}

// Create log Request godoc
// @Summary Create a new log request
// @Description Create a new log request and send them to kafka
// @Tags log
// @Accept  json
// @Produce  json
// @Param log body dtos.Log true "Create log"
// @Success 200 {object} dtos.ResStruct
// @Failure 400 {object} dtos.Res400Struct
// @Failure 500 {object} dtos.Res500Struct
// @Router /v1/logs [post]
func QueueLog(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rd := logAndGetContext(w, r)
	var status int
	defer func() {
		go prom.RecordHttpResponseTime(status, "Log", http.MethodPost, rd.Start)
	}()

	q, err := queue.New(rd.l)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		status = http.StatusInternalServerError
		return
	}

	decoder := json.NewDecoder(r.Body)
	req := &dtos.Log{}
	err = decoder.Decode(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusBadRequest, rd)
		status = http.StatusBadRequest
		return
	}

	err = q.QueueLog(req)
	if err != nil {
		writeJSONMessage(err.Error(), ERR_MSG, http.StatusInternalServerError, rd)
		status = http.StatusInternalServerError
		return
	}

	status = http.StatusOK
	writeJSONMessage("SUCCESS", MSG, http.StatusOK, rd)
}

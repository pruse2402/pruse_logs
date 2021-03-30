package config

import (
	"os"
	"pruse_logs/log-wrapper/internal/config/globals"
	"strings"

	"github.com/FenixAra/go-util/log"
)

var (
	APP_NAME      string
	PORT          string
	SERVICE_TOKEN string

	LOG_FILE_PATH_SIZE string
	LOG_LEVEL          string

	KAFKA_ADDR   []string
	GELF_VERSION string
	KAFKA_TOPIC  string
)

func init() {
	APP_NAME = os.Getenv("APP_NAME")
	PORT = os.Getenv("PORT")
	SERVICE_TOKEN = os.Getenv("SERVICE_TOKEN")

	LOG_LEVEL = os.Getenv("LOG_LEVEL")
	LOG_FILE_PATH_SIZE = os.Getenv("LOG_FILE_PATH_SIZE")

	KAFKA_ADDR = strings.Split(os.Getenv("KAFKA_ADDR"), ",")
	GELF_VERSION = os.Getenv("GELF_VERSION")
	KAFKA_TOPIC = os.Getenv("KAFKA_TOPIC")

	config := log.NewConfig("", LOG_LEVEL, LOG_FILE_PATH_SIZE, APP_NAME, "", "", "")
	globals.Logger = log.New(config)
}

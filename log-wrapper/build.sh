#!/bin/bash

export APP_NAME="LogWrapper"
export ENV="dev"
export DMS_URL=""
export PORT=9008
export LOG_LEVEL=Debug
export LOG_FILE_PATH_SIZE=Short
export KAFKA_ADDR=127.0.0.1:9092
export GELF_VERSION=1.1
export KAFKA_TOPIC=testing
export SERVICE_TOKEN=devtoken
 
# Sensitive environment variables are added in dev.env
source dev.env

swag init

go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi

echo "Restoring all vendor versions ..."
godep restore
echo "Done."


echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

echo "Running go vet ..."
go vet ./internal/...
if [ $? != 0 ]; then
  exit
fi
echo "Done."

echo "Running go generate ..."
go generate ./internal/...
echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

echo "Running unit test ..."
go test -p=1 ./internal/...
if [ $? == 0 ]; then
    echo "Done."
	echo "## Starting service ##"
    ./log-wrapper
fi

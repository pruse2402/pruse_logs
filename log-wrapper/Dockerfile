# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10

# Copy the local package files to the container's workspace.
ADD . /go/src/pruse_logs/log-wrapper

# Build the log-wrapper command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install pruse_logs/log-wrapper

# Run the log-wrapper command by default when the container starts.
ENTRYPOINT /go/bin/log-wrapper

# Document that the service listens on port 3000.
EXPOSE 3000

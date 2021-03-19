FROM golang:1.15.10-alpine3.13
# Set the Current Working Directory inside the container
RUN apk update
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git
ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:/usr/local/go/bin
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get github.com/golang/glog
RUN go get github.com/silkeh/senml

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

CMD ["./main"]


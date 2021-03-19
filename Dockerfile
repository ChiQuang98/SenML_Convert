FROM golang:lastest
WORKDIR /go/src/app
COPY . .
ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH $PATH:/usr/local/go/bin
RUN go get github.com/eclipse/paho.mqtt.golang
RUN go get github.com/golang/glog
RUN go get github.com/silkeh/senml



package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"SenML_cv/services"
	"os/signal"
	"syscall"
)
func init() {
	//glog
	//create logs folder
	os.Mkdir("./logs", 0777)
	flag.Lookup("stderrthreshold").Value.Set("[INFO|WARN|FATAL]")
	flag.Lookup("logtostderr").Value.Set("false")
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	glog.MaxSize = 1024 * 1024 * 258
	flag.Lookup("v").Value.Set(fmt.Sprintf("%d", 8))
	flag.Parse()

}
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	err:=services.ConnectMqtt()
	if err != nil{
		fmt.Println("Fail to connect to MQTT Broker")
		glog.Error("Failed err: ", err)
		os.Exit(1)
	}else {
		fmt.Println("OK MQTT Connect")
		glog.Info("OK")
	}
	glog.Flush()
	<-c
}

package services

import (
	"SenML_cv/base"
	"SenML_cv/settings"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/cisco/senml"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/glog"
	"github.com/krylovsk/gosenml"
	_ "github.com/krylovsk/gosenml"
	"strconv"
	"time"
)
var mqttInfo *settings.MqttInfo =settings.GetMqttInfo()
var mqttClient MQTT.Client
func convertJsonToSenML(mcu_id int64,obj interface{},op_code byte) (string,error){
	switch op_code {
	case base.OPU_GENERIC:
		fmt.Println("IN OPU_GENERIC")
		var generic *base.OPUGeneric = obj.(*base.OPUGeneric)
		volumn := float64(generic.Volume)
		timeLong:= (time.Now().Unix())
		entry :=gosenml.Entry{
			Name:         strconv.FormatInt(mcu_id, 10),
			Units:        "sound",
			Value:        &volumn,
			//StringValue:  &generic.FirmwareVersion,
			Time:         timeLong,
		}

		msg := gosenml.NewMessage(entry)
		err := msg.Validate()
		msg.BaseTime = timeLong
		msg.BaseName = "Generic"
		if err!=nil{
			return "",err
		}
		encoder := gosenml.NewJSONEncoder()
		rsString,_ := encoder.EncodeMessage(msg)
		fmt.Println(string(rsString))
		return string(rsString),nil

		break
	case base.OPU_SENSOR:
		fmt.Println("ConvertJsonToSenML_InSensor")
		var sensors []base.OPUSensor = obj.([]base.OPUSensor)
		fmt.Println(len(sensors))
		entries :=make([] gosenml.Entry,0)
		var m1 *gosenml.Message
		if len(sensors) == 0{
			return "", errors.New("Number of Array sensor is 0")
		}
		for _,sensor:=range sensors{
			v := float64(sensor.Value)
			e :=gosenml.Entry{
				Name:  strconv.FormatInt(mcu_id, 10),
				Units: "sensorU",
				Time: sensor.CreatedTime,
				Value: &v,
			}
			entries = append(entries,e)
			fmt.Println(sensor.Name)
		}
		m1=gosenml.NewMessage(entries...)
		m1.BaseName = "http://example.com/"
		err := m1.Validate()
		//m1.Version = 2
		if err != nil{
			return "",err
			fmt.Println("err",err.Error())
		}
		encoder := gosenml.NewJSONEncoder()
		b, _ := encoder.EncodeMessage(m1)
		fmt.Println(string(b))
		return string(b), nil
		break
	}
	return "",errors.New("Wrong ")
}
func ConnectMqtt() error {
	defer func() {
		if err := recover(); err != nil {
			glog.Error("-------------RECOVER err: ", err)
		}
	}()
	server := fmt.Sprintf("tcp://%s:%d", mqttInfo.ServerAddress, mqttInfo.ServerPort)
	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID("Test")
	opts.SetCleanSession(false)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetAutoReconnect(true)
	opts.SetConnectionLostHandler(onConnectionLost)
	opts.SetOnConnectHandler(onConnected)
	mqttClient = MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
var onConnected MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("connected")
	if token := client.Subscribe("resultTopic",1,onMessage); token.Wait() && token.Error()!=nil{
		glog.Error("onConnected/client.Subscribe(resultTopic) err: ", token.Error())
		return
	}
}
var onConnectionLost MQTT.ConnectionLostHandler = func(client MQTT.Client, reason error) {
	glog.Infof("onConnectionLost/server(%s:%d)", mqttInfo.ServerAddress, mqttInfo.ServerPort)
}
var onMessage MQTT.MessageHandler= func(client MQTT.Client, msg MQTT.Message) {
	fmt.Println("INNN")
	obj := &base.MqttMessagePublish{}
	err := json.Unmarshal(msg.Payload(), obj)
	var MCUID int64 = obj.Id
	if err != nil {
		glog.Error("onMessage/json.Unmarshal err: ", err)
		glog.Error("Raw message: ", msg.Payload())
		return
	}
	switch obj.OpCode {
	case base.OPU_GENERIC:

		generic := &base.OPUGeneric{}
		err := json.Unmarshal(obj.Data, generic)
		if err != nil {
			glog.Errorf("onMessage/OPU_GENERIC/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		convertJsonToSenML(MCUID,generic,base.OPU_GENERIC)
		break
	case base.OPU_CAMERA:
		cameras := []base.OPUCamera{}
		err := json.Unmarshal(obj.Data, &cameras)
		if err != nil {
			glog.Errorf("onMessage/OPU_CAMERA/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_PHONE:
		phone := ""
		err := json.Unmarshal(obj.Data, &phone)
		if err != nil {
			glog.Errorf("onMessage/OPU_PHONE/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_SENSOR:
		fmt.Println("IN SENSOR")
		sensors := []base.OPUSensor{}
		err := json.Unmarshal(obj.Data, &sensors)
		if err != nil {
			glog.Errorf("onMessage/OPU_SENSOR/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		convertJsonToSenML(MCUID,sensors,base.OPU_SENSOR)

		break
	case base.OPU_ALARM:
		alarms := []base.OPUAlarm{}
		err := json.Unmarshal(obj.Data, &alarms)
		if err != nil {
			glog.Errorf("onMessage/OPU_ALARM/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}

		break
	case base.OPU_STATUS:
		status := base.OPUStatus{}
		err := json.Unmarshal(obj.Data, &status)
		if err != nil {
			glog.Errorf("onMessage/OPU_STATUS/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_MEDIA:
		medias := [][]int64{}
		err := json.Unmarshal(obj.Data, &medias)
		if err != nil {
			glog.Errorf("onMessage/OPU_MEDIA/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		break
	case base.OPU_LOG:
		logs := []base.OPULog{}
		err := json.Unmarshal(obj.Data, &logs)
		if err != nil {
			glog.Errorf("onMessage/OPU_LOG/%d/json.Unmarshal err: %v", obj.Id, err)
			return
		}
		break
	}
}

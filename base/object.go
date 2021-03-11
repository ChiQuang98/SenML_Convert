package base

import "encoding/json"

type MqttMessagePublish struct {
	Id     int64           `json:"id"`
	OpCode byte            `json:"opcode"`
	Data   json.RawMessage `json:"data"`
}
type OPUGeneric struct {
	GroupList       []int64     `json:"group"`
	MediaIdLastest  int64       `json:"mid"`
	Volume          byte        `json:"volume"`
	LocalIp         string      `json:"ip"`
	VCode           string      `json:"vcode"`
	PhoneNumber     string      `json:"phone"`
	CameraList      []OPUCamera `json:"camera"`
	SensorList      []OPUSensor `json:"sensor"`
	AlarmList       []OPUAlarm  `json:"alarm"`
	FirmwareVersion string      `json:"fvers"`
	FMVolume        int32       `json:"fmvolume"`
	FMAuto          int32       `json:"fmauto"`
	TxType          int32       `json:"txtype"`
	ConnStatus      byte
	ConnTime        int64
	WanIP           string
}
type OPUCamera struct {
	CameraId       int64  `json:"id"`
	CameraName     string `json:"name"`
	CameraLocalIp  string `json:"ip"`
	CameraTypeId   int32  `json:"type"`
	CameraHttpPort int32  `json:"http"`
	CameraRtspPort int32  `json:"rtsp"`
	CameraUsername string `json:"user"`
	CameraPassword string `json:"pass"`
}

type OPUPhoneNumber struct {
	PhoneNumber string `json:"phone_number"`
}

type OPUSensor struct {
	SensorId    int64   `json:"id"`
	Enable      byte    `json:"enable"`
	Name        string  `json:"name"`
	Type        int32   `json:"type"`
	State       byte    `json:"state"`
	Value       int64   `json:"value"`
	Thresholds  []int32 `json:"threshold"`
	CreatedTime int64   `json:"ctime"`
}

type OPUAlarm struct {
	EventId      int64  `json:"id"`
	EventType    int32  `json:"type"`
	Name         string `json:"name"`
	State        byte   `json:"state"`
	SensorId     int64  `json:"sid"`
	SensorState  byte   `json:"alarm"`
	Mode         byte   `json:"mode"`
	ActiveTime   int16  `json:"active"`
	InactiveTime int16  `json:"inactive"`
	AutoDays     byte   `json:"days"`
	PlayFile     int64  `json:"mid"`
	OccurTime    int64  `json:"occur"`
}
type OPULog struct {
	LogId       int64 `json:"id"`
	CreatedTime int64 `json:"time"`
	LogType     byte  `json:"type"`
	MediaId     int64 `json:"mid"`
	EventId     int64 `json:"eid"`
	SensorId    int64 `json:"sid"`
	State       byte  `json:"state"`
	Value       int64 `json:"value"`
}
type OPUStatus struct {
	TxType     byte   `json:"conn"`
	Temp       int16  `json:"temp"`
	SpeakerErr byte   `json:"spkerr"`
	SpeakerSta uint16 `json:"spksta"`
	MCsq       byte   `json:"csqm"`
	WiCsq      byte   `json:"csqw"`
	FMStatus   byte   `json:"fmsta"`
}
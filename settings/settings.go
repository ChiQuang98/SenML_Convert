package settings

import (
	"encoding/json"
	"io/ioutil"
)

var settings Settings = Settings{}

func init()  {
	content, err := ioutil.ReadFile("setting.json")
	if err != nil {
		panic(err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		panic(jsonErr)
	}
}
type Settings struct{
	MqttInfo             *MqttInfo
	KeyAuth				 string
}
type MqttInfo struct{
	ServerAddress   string
	ServerPort      int
	HttpApiPort     int
}
func GetMqttInfo() *MqttInfo {
	return settings.MqttInfo
}
package main

import (
	"fmt"
	"time"

	"github.com/silkeh/senml"
)

func main() {
	now := time.Now()
	list := []senml.Measurement{
		senml.NewValue("sensor:temperature", 23.5, senml.Celsius, now, 0),
		senml.NewValue("sensor:humidity", 33.7, senml.RelativeHumidityPercent, now, 0),
	}
	list = append(list, senml.NewValue("sensor:humidity11", 40, "sound", now, 0))
	fmt.Print(len(list))
	data, err := senml.EncodeJSON(list)
	if err != nil {
		fmt.Print("Error encoding to JSON:", err)
	}

	fmt.Printf("%s\n", data)
}
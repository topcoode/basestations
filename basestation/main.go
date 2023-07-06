package main

import (
	"basestation/signal"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Basestation struct {
	BTNAME   string `json:"basestationname"`
	BTID     int    `json:"basestationID"`
	Signal   int    `json:"signal"`
	length   int    `json:"length"`
	minValue int    `json:"minvalue"`
	maxValue int    `json:"maxvalue"`
}
type Signal struct {
	Signaldata float64
}

func BT(c *gin.Context) {
	fmt.Println("the value of c: ", c.Request.Body)
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("*******", string(body))
	var data Signal
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("data.signal-->", data.Signaldata)
	jsondata, err := json.MarshalIndent(c, "", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsondata))
	basestation := []Basestation{
		{
			BTNAME:   "basestation1",
			BTID:     0000000000001,
			Signal:   50,
			length:   10,
			minValue: 0,
			maxValue: 20,
		},
		{
			BTNAME:   "basestation2",
			BTID:     0000000000001,
			Signal:   50,
			length:   10,
			minValue: 0,
			maxValue: 20,
		},
		{
			BTNAME:   "basestation3",
			BTID:     0000000000001,
			Signal:   50,
			length:   10,
			minValue: 0,
			maxValue: 20,
		},
		{
			BTNAME:   "basestation4",
			BTID:     0000000000001,
			Signal:   50,
			length:   10,
			minValue: 0,
			maxValue: 20,
		},
	}

	signal.GenerateRandomSignal()
	basestationdata, _ := json.MarshalIndent(basestation, "", "")
	fmt.Println(string(basestationdata))
}
func main() {
	router := gin.Default()
	router.POST("/basestation", BT)
	router.Run(":8081")
}

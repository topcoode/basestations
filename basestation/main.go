package main

import (
	"basestation/signal"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Basestation struct {
	BTNAME       string  `json:"basestationname"`
	BTID         float64 `json:"basestationID"`
	Ueconnection *Ueconnectiondata
}

type Ueconnectiondata struct {
	Uedatasignal      *Uedata
	Basestationsignal float64 `json:"basestationsignal"`
	Ping              float64
}
type Uedata struct {
	Uesignal float64 `json:"uesignal"`
}

func BT(c *gin.Context) {
	var data Basestation
	fmt.Println("client data----->", data)
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	body, _ := ioutil.ReadAll(c.Request.Body)
	var Ping float64 //Produce
	var Uesignal float64
	var ServerSignal float64 //produce
	fmt.Println(ServerSignal)
	basestation := []Basestation{
		{
			BTNAME: "Basestation1",
			BTID:   0000000000001,
			Ueconnection: &Ueconnectiondata{
				Uedatasignal: &Uedata{
					Uesignal: Uesignal,
				},
				Basestationsignal: 100,
				Ping:              Ping,
			},
		},
		{
			BTNAME: "Basestation2",
			BTID:   0000000000002,
			Ueconnection: &Ueconnectiondata{
				Uedatasignal: &Uedata{
					Uesignal: Uesignal,
				},
				Basestationsignal: 100,
				Ping:              Ping,
			},
		},
	}
	// fmt.Println("ue signal---->", data.Ueconnection.Uedatasignal.Uesignal)
	for _, user := range basestation {

		//parse json----------------------------------------------->
		var data Basestation
		if err := json.Unmarshal(body, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Data received successfully"})
		//server signal---------->
		serversignal := signal.GenerateRandomSignal()
		fmt.Println("SERVER SIGNAL:", serversignal)
		// ue signal------------>
		user.Ueconnection.Ping = serversignal - user.Ueconnection.Basestationsignal
		Ping = user.Ueconnection.Ping
		Uesignal = user.Ueconnection.Uedatasignal.Uesignal
		ServerSignal = serversignal
		fmt.Println("user--------->", user)
	}
}

func main() {
	router := gin.Default()
	router.POST("/basestation", BT)
	router.Run(":8081")
}

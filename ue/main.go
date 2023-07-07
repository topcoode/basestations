package main

import (
	"basestation/signal"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Basestation struct {
	BTNAME       string  `json:"basestationname"`
	BTID         float64 `json:"basestationID"`
	Ueconnection Ueconnectiondata
}

type Ueconnectiondata struct {
	Uedatasignal      Uedata
	Basestationsignal float64 `json:"basestationsignal"`
	Ping              float64
}
type Uedata struct {
	Uesignal float64 `json:"uesignal"`
}

func Ueconnection(c *gin.Context) {
	var data Basestation
	fmt.Println("server data-------->", data)
	UESignalgenerated := signal.GenerateRandomSignal()

	uesignal := Uedata{
		Uesignal: UESignalgenerated,
	}
	fmt.Println("uesignal:--------->", uesignal)
	jsondata, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsondata))

	req, err := http.NewRequest("POST", "http://localhost:8081/basestation", bytes.NewBuffer(jsondata))
	fmt.Println(req)
	if err != nil {
		fmt.Println("error in link", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
func main() {
	router := gin.Default()
	router.POST("/ue", Ueconnection)
	router.Run(":8080")
}

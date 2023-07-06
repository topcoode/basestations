package main

import (
	"basestation/signal"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ue struct {
	UeId     int
	Signal   int
	length   int
	minValue int
	maxValue int
}
type Signal struct {
	Signaldata float64
}

func Ueconnection(c *gin.Context) {
	ue := Ue{
		UeId:     100000000000,
		Signal:   50,
		length:   10,
		minValue: 0,
		maxValue: 20,
	}
	fmt.Println(ue)
	signal := signal.GenerateRandomSignal()
	fmt.Println("UE SIGNAL...", signal)
	// uedata, _ := json.MarshalIndent(ue, "", "")
	// fmt.Println("uedata :", uedata)
	jsondata, err := json.Marshal(signal)
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

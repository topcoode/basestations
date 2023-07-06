package signal

import (
	"math/rand"
	"time"
)

type Signal struct {
	Signaldata float64
}

func GenerateRandomSignal() float64 { //(length, minValue, maxValue int) []int {
	var signal Signal
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64()
	signal.Signaldata = randomNumber
	return randomNumber
}

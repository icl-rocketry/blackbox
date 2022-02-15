package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var dataChannel = make(chan data)

type tripleSensor struct {
	X float32
	Y float32
	Z float32
}

type accelerometer struct {
	tripleSensor
}

type gyroscope struct {
	tripleSensor
}

type magnetometer struct {
	tripleSensor
}

type data struct {
	Time         int
	Acceleration accelerometer
	Gyroscope    gyroscope
	Magnetometer magnetometer
	Pressure     float32
	Temperature  float32
}

func dataHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Thanos won")
	}

	for d := range dataChannel {
		if err = conn.WriteJSON(d); err != nil {
			return
		}
	}
}

func main() {
	r := gin.Default()

	go func() {
		d := data{
			Time: 3,
			Acceleration: accelerometer{
				tripleSensor: tripleSensor{X: 1, Y: 2, Z: 3},
			},
			Gyroscope: gyroscope{
				tripleSensor: tripleSensor{X: 1, Y: 2, Z: 3},
			},
			Magnetometer: magnetometer{
				tripleSensor: tripleSensor{X: 1, Y: 2, Z: 3},
			},
			Pressure:    1.02,
			Temperature: 32.1,
		}
		t := time.NewTicker(time.Second)
		for range t.C {
			dataChannel <- d
		}
	}()

	r.GET("/ws", dataHandler)
	r.Use(static.Serve("/", static.LocalFile("../blackbox-frontend/dist", true)))

	r.Run()
}

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

var (
	recording    = false
	recordedData = make([]data, 0)
	saveChannel  = make(chan struct{}) // Use to save data
	dataChannel  = make(chan data)
)

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

type location struct {
	Latitude  float32
	Longitude float32
}

type data struct {
	Time         int
	Acceleration accelerometer
	Gyroscope    gyroscope
	Magnetometer magnetometer
	Pressure     float32
	Temperature  float32
	Location     location
}

func startHandler(ctx *gin.Context) {
	if recording {
		ctx.JSON(http.StatusBadRequest, "Recording already started")
	} else {
		recording = true
		ctx.JSON(http.StatusAccepted, "Recording started")
	}
}

func endHandler(ctx *gin.Context) {
	if recording {
		recording = false
		ctx.JSON(http.StatusAccepted, "Recording ended")
		saveChannel <- struct{}{}
	} else {
		ctx.JSON(http.StatusBadRequest, "Not currently recording")
	}
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

func flush() {
	for range saveChannel {
		files, err := ioutil.ReadDir("./data/")
		if err != nil {
			log.Println("Couldn't read data directory because", err)
			continue
		}
		filename := fmt.Sprintf("data_%d.csv", len(files))
		file, err := os.Create("./data/" + filename)
		if err != nil {
			log.Println("Couldn't create file because", err)
			continue
		}
		w := csv.NewWriter(file)
		w.Write([]string{
			"Time",
			"Acc_x", "Acc_y", "Acc_z",
			"Gyro_x", "Gyro_y", "Gyro_z",
			"Mag_x", "Mag_y", "Mag_z",
			"Pressue", "Temperature",
			"Longitude", "Latitude",
		})

		for _, d := range recordedData {
			err = w.Write([]string{
				fmt.Sprintf("%d, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f",
					d.Time,
					d.Acceleration.X, d.Acceleration.Y, d.Acceleration.Z,
					d.Gyroscope.X, d.Gyroscope.Y, d.Gyroscope.Z,
					d.Magnetometer.X, d.Magnetometer.Y, d.Magnetometer.Z,
					d.Pressure, d.Temperature,
					d.Location.Longitude, d.Location.Latitude),
			})
			if err != nil {
				log.Println("Couldn't log row because", err)
				continue
			}
		}
		w.Flush()
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
			Location: location{
				Latitude:  51.5007,
				Longitude: -0.1246,
			},
		}
		t := time.NewTicker(time.Second)
		for range t.C {
			dataChannel <- d
		}
	}()

	go flush()

	r.GET("/ws", dataHandler)
	r.POST("/start", startHandler)
	r.POST("/end", endHandler)
	r.Use(static.Serve("/", static.LocalFile("../blackbox-frontend/dist", true)))
	r.Use(static.Serve("/public/", static.LocalFile("../blackbox-frontend/public", true)))

	r.Run()
}

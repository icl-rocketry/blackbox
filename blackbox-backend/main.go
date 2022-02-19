package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const DATA_SIZE = 4 + 12 + 12 + 12 + 4 + 4 + 8

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
	Time         int32
	Acceleration accelerometer
	Gyroscope    gyroscope
	Magnetometer magnetometer
	Pressure     float32
	Temperature  float32
	Location     location
}

func fromBuf(buf *bytes.Buffer) data {
	return data{
		Time: int32(binary.BigEndian.Uint32(buf.Next(4))),
		Acceleration: accelerometer{
			tripleSensor: tripleSensor{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},
		Gyroscope: gyroscope{
			tripleSensor: tripleSensor{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},
		Magnetometer: magnetometer{
			tripleSensor: tripleSensor{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},
		Pressure:    math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
		Temperature: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
		Location: location{
			Latitude:  math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			Longitude: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
		},
	}
}

func startHandler(ctx *gin.Context) {
	if recording {
		ctx.JSON(http.StatusBadRequest, "Recording already started")
	} else {
		recording = true
		ctx.JSON(http.StatusOK, "Recording started")
	}
}

func endHandler(ctx *gin.Context) {
	if recording {
		recording = false
		ctx.JSON(http.StatusOK, "Recording ended")
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
		defer file.Close()
		file.Write([]byte("Time, Acc_x, Acc_y, Acc_z, Gyro_x, Gyro_y, Gyro_z, Mag_x, Mag_y, Mag_z, Pressue, Temperature, Longitude, Latitude"))

		for _, d := range recordedData {
			_, err = file.Write([]byte(
				fmt.Sprintf("%d, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f",
					d.Time,
					d.Acceleration.X, d.Acceleration.Y, d.Acceleration.Z,
					d.Gyroscope.X, d.Gyroscope.Y, d.Gyroscope.Z,
					d.Magnetometer.X, d.Magnetometer.Y, d.Magnetometer.Z,
					d.Pressure, d.Temperature,
					d.Location.Longitude, d.Location.Latitude),
			))
			if err != nil {
				log.Println("Couldn't log row because", err)
				continue
			}
		}
	}
}

func udpServer() {
	pc, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, DATA_SIZE*50)
		n, _, err := pc.ReadFrom(buf)
		if err != nil || n < DATA_SIZE || !recording {
			log.Println("Couldn't read from UDP socket because", err, n, recording)
			continue
		}

		go serve(bytes.NewBuffer(buf[:n]))
	}
}

func serve(buf *bytes.Buffer) {
	data := fromBuf(buf)
	dataChannel <- data
	recordedData = append(recordedData, data)
}

func main() {
	r := gin.Default()

	go flush()
	go udpServer()

	r.GET("/ws", dataHandler)
	r.POST("/start", startHandler)
	r.POST("/end", endHandler)
	r.Use(static.Serve("/", static.LocalFile("../blackbox-frontend/dist", true)))
	r.Use(static.Serve("/public/", static.LocalFile("../blackbox-frontend/public", true)))

	r.Run()
}

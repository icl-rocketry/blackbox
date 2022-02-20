package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

const DATA_SIZE = 4 + 12 + 12 + 12 + 4 + 4 + 8

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

package main

import (
	"bytes"
	"encoding/binary"
	"math"
)

const DATA_SIZE = 4 + 12 + 12 + 12 + 12 + 12 + 4 + 4 + 4

type tripleCartesian struct {
	X float32
	Y float32
	Z float32
}

type orientation struct {
	tripleCartesian
}

type accelerometer struct {
	tripleCartesian
}

type gyroscope struct {
	tripleCartesian
}

type magnetometer struct {
	tripleCartesian
}

type location struct {
	Latitude  float32
	Longitude float32
	altitude float32
}

type data struct {
	Time         int32
	Orientation  orientation
	Gyroscope    gyroscope
	Acceleration accelerometer	
	Magnetometer magnetometer
	Location     location
	Temperature  float32
	Pressure     float32
	Lux 		 float32
	
}

func fromBuf(buf *bytes.Buffer) data {
	return data{
		Time: int32(binary.BigEndian.Uint32(buf.Next(4))),
		Orientation: orientation{
			tripleCartesian: tripleCartesian{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},
		Gyroscope: gyroscope{
			tripleCartesian: tripleCartesian{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},
		Acceleration: accelerometer{
			tripleCartesian: tripleCartesian{
				X: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Y: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
				Z: math.Float32frombits(binary.BigEndian.Uint32(buf.Next(4))),
			},
		},		
		Magnetometer: magnetometer{
			tripleCartesian: tripleCartesian{
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

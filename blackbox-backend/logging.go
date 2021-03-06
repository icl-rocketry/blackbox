package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func flush() {
	for range saveChannel {
		files, err := ioutil.ReadDir("./data/")
		if err != nil {
			log.Println("Creating data directory because:", err)
			err = os.Mkdir("./data/", os.ModePerm)
			if err != nil {
				log.Println("Couldn't create data directory because: ", err)
				continue
			}
		}
		filename := fmt.Sprintf("data_%d.csv", len(files))
		file, err := os.Create("./data/" + filename)
		if err != nil {
			log.Println("Couldn't create file because", err)
			continue
		}
		file.Write([]byte("Time, Or_x, Or_y, Or_z, Gyro_x, Gyro_y, Gyro_z, Acc_x, Acc_y, Acc_z, Mag_x, Mag_y, Mag_z, Latitude, Longitude, Altitude, Temperature, Lux, Pressure"))

		for _, d := range recordedData {
			_, err = file.Write([]byte(
				fmt.Sprintf("%d, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f, %f",
					d.Time,
					d.Orientation.X, d.Orientation.Y, d.Orientation.Z,
					d.Gyroscope.X, d.Gyroscope.Y, d.Gyroscope.Z,
					d.Acceleration.X, d.Acceleration.Y, d.Acceleration.Z,
					d.Magnetometer.X, d.Magnetometer.Y, d.Magnetometer.Z,
					d.Location.Longitude, d.Location.Latitude, d.Location.Altitude,
					d.Temperature, d.Lux, d.Pressure),
			))
			if err != nil {
				log.Println("Couldn't log row because", err)
				continue
			}
		}
		file.Close()
	}
}

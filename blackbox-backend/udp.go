package main

import (
	"bytes"
	"log"
	"net"
)

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

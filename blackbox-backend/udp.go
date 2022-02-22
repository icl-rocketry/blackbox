package main

import (
	"bytes"
	"log"
	"net"
)

const UDP_PORT = "2052"

func udpServer() {
	pc, err := net.ListenPacket("udp", ":"+UDP_PORT)
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
	clientMutex.RLock()
	for i, dataChannel := range clients {
		log.Println("Sending data to client: ", i)
		dataChannel <- data
	}
	clientMutex.RUnlock()
	recordedData = append(recordedData, data)
}

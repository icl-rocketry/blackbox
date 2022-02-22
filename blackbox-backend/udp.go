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
	recvd := fromBuf(buf)
	clientMutex.RLock()
	defer clientMutex.RUnlock()
	for i, dataChannel := range clients {
		// This is terrible - Akshat Tripathi Tuesday 22/2/22
		go func(dataChannel chan<- data, i int) {
			defer func(i int) {
				clientMutex.Lock()
				defer clientMutex.Unlock()
				if r := recover(); r != nil {
					log.Println("Recovered in f", r)
					clients = append(clients[:i], clients[i+1:]...)
				}
			}(i)

			log.Println("Sending data to client: ", i)
			dataChannel <- recvd
		}(dataChannel, i)
	}
	recordedData = append(recordedData, recvd)
}

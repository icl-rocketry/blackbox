package main

/*This could be better abstracted, but since we aren't really using it too heavily, this is alright
as an MVP.

A client connects to ws://localhost:PORT/ws, which registers a channel in clients.
When a udp packet arrives (from anyone in the world - nice security issue there), if we're recording,
then record it and broadcast it to all clients.
When a client receives from its channel, it sends the data via websocket to the frontend it's connected
to.

To start recording, send an empty POST request to /start.
To end recording, send an empty POST request to /end.

*/

import (
	"log"
	"net/http"
	"os"
	"sync"

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
	clientMutex  = sync.RWMutex{}
	clients      = make([]chan<- data, 0)
)

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
		ctx.JSON(http.StatusInternalServerError, "Couldn't upgrade to websocket")
	}

	dataChannel := make(chan data)
	addClient(dataChannel)

	for d := range dataChannel {
		if err = conn.WriteJSON(d); err != nil {
			return // TODO: This causes a memory leak but MVP
		}
	}
}

func addClient(channel chan<- data) {
	clientMutex.Lock()
	clients = append(clients, channel)
	clientMutex.Unlock()
}

func main() {
	HTML_PATH := os.Getenv("HTML_PATH")
	log.Println(HTML_PATH)
	r := gin.Default()

	go flush()
	go udpServer()

	r.GET("/ws", dataHandler)
	r.POST("/start", startHandler)
	r.POST("/end", endHandler)
	r.Use(static.Serve("/", static.LocalFile(HTML_PATH+"/dist", true)))
	r.Use(static.Serve("/public/", static.LocalFile(HTML_PATH+"/public", true)))

	r.Run()
}

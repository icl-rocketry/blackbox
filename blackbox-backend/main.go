package main

import (
	"net/http"

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

	for d := range dataChannel {
		if err = conn.WriteJSON(d); err != nil {
			return
		}
	}
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

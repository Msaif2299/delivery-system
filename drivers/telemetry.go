package drivers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type TelemetryRequest struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// TODO: add protection against CSRF attacks
		return true
	},
}

func UpdateTelemetryData(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Printf("Error encountered in UpdateTelemetryData, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Could not upgrade to websocket"})
		return
	}
	defer conn.Close()
	for {
		mtype, msg, err := conn.ReadMessage()
		if websocket.IsCloseError(err, websocket.CloseGoingAway) {
			break
		}
		if err != nil {
			fmt.Printf("Error in reading message in UpdateTelemetryData, mtype: %d, msg: %s, err: %s\n", mtype, msg, err.Error())
			break
		}
		var req TelemetryRequest
		if err := json.Unmarshal(msg, &req); err != nil {
			fmt.Printf("Error in unmarshaling the data, received: %s err: %s\n", string(msg), err.Error())
			continue
		}
		fmt.Printf("Received message: %+v\n", req)
	}
}

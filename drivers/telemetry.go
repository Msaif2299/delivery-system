package drivers

import (
	"delivery-system/datastore"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type TelemetryRequest struct {
	VehicleLicensePlate string  `json:"license_plate"`
	Latitude            string  `json:"latitude,omitempty"`
	Longitude           string  `json:"longitude,omitempty"`
	Speed               float64 `json:"speed,omitempty"`
	EngineTemp          int     `json:"engine_temp,omitempty"`
	TirePressure        struct {
		FrontLeft  int32 `json:"front_left,omitempty"`
		FrontRight int32 `json:"front_right,omitempty"`
		BackLeft   int32 `json:"back_left,omitempty"`
		BackRight  int32 `json:"back_right,omitempty"`
	} `json:"tire_pressure,omitempty"`
}

func (req TelemetryRequest) ConvertToDTO() map[string]interface{} {
	return NewTelemetryDTO(
		req.Latitude,
		req.Longitude,
		req.Speed,
		req.EngineTemp,
		req.TirePressure.FrontLeft,
		req.TirePressure.FrontRight,
		req.TirePressure.BackLeft,
		req.TirePressure.FrontRight,
	)
}

func NewTelemetryDTO(
	latitude string,
	longitude string,
	speed float64,
	engineTemp int,
	tirePressureFrontLeft int32,
	tirePressureFrontRight int32,
	tirePressureBackLeft int32,
	tirePressureBackRight int32,
) map[string]interface{} {
	return map[string]interface{}{
		"latitude":                  latitude,
		"longitude":                 longitude,
		"speed":                     speed,
		"engine_temp":               engineTemp,
		"tire_pressure_front_left":  tirePressureFrontLeft,
		"tire_pressure_front_right": tirePressureFrontRight,
		"tire_pressure_back_left":   tirePressureBackLeft,
		"tire_pressure_back_right":  tirePressureBackRight,
	}
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
	db := datastore.GetNoSQLDataStore(c)
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
		if req.VehicleLicensePlate == "" {
			fmt.Println("Error in UpdateTelemetryData, license plate not found, discarding message...")
			continue
		}
		db.WriteAsync(c.Request.Context(), req.VehicleLicensePlate, map[string]string{}, req.ConvertToDTO())
		fmt.Printf("Received message: %+v\n", req)
	}
}

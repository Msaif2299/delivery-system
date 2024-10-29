package vehicles

import (
	"delivery-system/datastore"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VehicleStatus uint8

const (
	VEHICLE_STATUS_IDLE        VehicleStatus = iota // If no location or speed data is received for a specified period, the vehicle’s status could automatically switch to Idle.
	VEHICLE_STATUS_IN_TRANSIT                       // When a vehicle begins moving at a defined speed threshold, its status might automatically update to In-Transit.
	VEHICLE_STATUS_MAINTENANCE                      // The status could change to Maintenance automatically when certain telemetry metrics indicate an issue (e.g., low tire pressure or high engine temperature).
	VEHICLE_STATUS_LOADING                          // Vehicle is loading at a dock.
	VEHICLE_STATUS_UNLOADING                        // Vehicle is unloading at a dock.
	VEHICLE_STATUS_RESERVED                         // The vehicle is reserved for a specific task or trip, but the trip hasn’t started yet.
	VEHICLE_STATUS_OFFLINE                          // If the telemetry feed stops for a significant time, the status could automatically switch to Offline.
)

const (
	VEHICLE_STATUS_MEASUREMENT string = "vehicle_status"
)

type SetStatusRequest struct {
	LicensePlate string `json:"license_plate"`
	Status       uint8  `json:"status"`
}

func ValidateVehicleStatus(status uint8) error {
	if status > uint8(VEHICLE_STATUS_OFFLINE) {
		return fmt.Errorf("invalid status, found: %d", status)
	}
	return nil
}

func SetVehicleStatusHandler(c *gin.Context) {
	var req SetStatusRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Printf("Error while binding json in SetVehicleStatusHandler, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid body"})
		return
	}
	if err := ValidateLicensePlateNumber(req.LicensePlate); err != nil {
		fmt.Printf("Error in SetVehicleStatusHandler, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid license plate"})
		return
	}
	if err := ValidateVehicleStatus(req.Status); err != nil {
		fmt.Printf("Error in SetVehicleStatusHandler, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid vehicle status"})
		return
	}
	db := datastore.GetNoSQLDataStore(c)
	db.WriteAsync(c.Request.Context(), VEHICLE_STATUS_MEASUREMENT, map[string]string{
		"license_plate": req.LicensePlate,
	}, map[string]interface{}{
		"status": req.Status,
	})
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

type GetStatusResponse struct {
	Status uint8 `json:"status"`
}

func GetVehicleStatusHandler(c *gin.Context) {
	licensePlate := c.Param("license_plate")
	if err := ValidateLicensePlateNumber(licensePlate); err != nil {
		fmt.Printf("Error in GetVehicleStatusHandler, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid license plate"})
		return
	}
	db := datastore.GetNoSQLDataStore(c)
	status, err := db.GetLastValue(c.Request.Context(), VEHICLE_STATUS_MEASUREMENT, map[string]string{
		"license_plate": licensePlate,
	})
	if err != nil {
		fmt.Printf("Error in GetVehicleStatusHandler, err: %s\n", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	res := status.(uint8)
	c.JSON(http.StatusOK, GetStatusResponse{Status: res})
}

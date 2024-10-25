package main

import (
	"delivery-system/datastore"
	"delivery-system/drivers"
	"delivery-system/vehicles"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var sqlDB datastore.SQLDataStore
	var err error
	sqlDB, err = datastore.NewMySQLDataStore()
	if err != nil {
		fmt.Printf("Could not establish connection, err: %s exiting...", err.Error())
		return
	}
	noSQLDB := datastore.NewInfluxDB2()

	r := gin.New()
	group := r.Group("api/v1/", datastore.SQLDBProvider(sqlDB))
	{
		group.GET("/drivers/fetch/:license_number", drivers.GetDriverHandler)
		group.GET("/vehicles/fetch/:license_plate", vehicles.GetVehicleHandler)

		group.POST("/drivers/register", drivers.RegisterDriverHandler)
		group.POST("/vehicles/register", vehicles.RegisterVehicleHandler)
		group.POST("/vehicles/assigndriver", vehicles.AssignDriverToVehicleHandler)

		group.PUT("/drivers/update", drivers.UpdateDriverInfoHandler)
		group.PUT("/vehicles/update", vehicles.UpdateVehicleInfoHandler)

		group.DELETE("/drivers/remove/:license_number", drivers.DeleteDriverInfoHandler)
		group.DELETE("/vehicles/remove/:license_plate", vehicles.DeleteVehicleInfoHandler)
	}

	websocketGroup := r.Group("websocket/v1/", datastore.NoSQLDBProvider(noSQLDB))
	{
		websocketGroup.GET("/vehicles/telemetry/update", drivers.UpdateTelemetryData)
	}

	r.Run(":8080")
}

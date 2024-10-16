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

	r := gin.New()
	group := r.Group("api/v1/", datastore.SQLDBProvider(sqlDB))
	{
		group.GET("/drivers/fetch/:id", drivers.GetDriverHandler)
		group.GET("/vehicles/fetch/:id", vehicles.GetVehicleHandler)

		group.POST("/drivers/register", drivers.RegisterDriverHandler)
		group.POST("/vehicles/register", vehicles.RegisterVehicleHandler)

		group.PUT("/drivers/update", drivers.UpdateDriverInfoHandler)
		group.PUT("/vehicles/update", vehicles.UpdateVehicleInfoHandler)

		group.DELETE("/drivers/remove", drivers.DeleteDriverInfoHandler)
		group.DELETE("/vehicles/remove", vehicles.DeleteVehicleInfoHandler)
	}

	r.Run(":8080")
}

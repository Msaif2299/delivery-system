package main

import (
	"delivery-system/datastore"
	"delivery-system/drivers"
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
		group.GET("/drivers/:id", drivers.GetDriverHandler)
	}

	r.Run(":7850")
}
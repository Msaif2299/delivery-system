package vehicles

import (
	"database/sql"
	"delivery-system/datastore"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetVehicleHandler(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no id param found"})
		return
	}
	query := "SELECT * FROM vehicles WHERE id = ?"
	db := datastore.GetSQLDataStore(c)
	var vehicle VehicleDTO
	row := db.QueryRow(query, ID)
	if err := row.Scan(&vehicle.ID, &vehicle.LicensePlate, &vehicle.Type, &vehicle.Make, &vehicle.Model, &vehicle.Year, &vehicle.CapacityKg, &vehicle.DriverID); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "vehicle id does not exist"})
			return
		}
		fmt.Printf("SQL Error in GetVehicleHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal error while searching"})
		return
	}
	c.JSON(http.StatusOK, ConvertVehicleDTOToRequest(vehicle))
}

func RegisterVehicleHandler(c *gin.Context) {
	var newVehicleRequest VehicleRequest
	if err := c.BindJSON(&newVehicleRequest); err != nil {
		fmt.Printf("Error binding POST body in RegisterVehicleHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}

	newVehicle := ConvertVehicleRequestToDTO(newVehicleRequest)
	query := "INSERT INTO vehicles (id, license_plate, type, make, model, year, capacity_kg, driver_id) VALUES (?,?,?,?,?,?,?,?)"
	db := datastore.GetSQLDataStore(c)
	id, err := db.Exec(query, newVehicle.ID, newVehicle.LicensePlate, newVehicle.Type, newVehicle.Make, newVehicle.Model, newVehicle.Year, newVehicle.CapacityKg, newVehicle.DriverID)
	if err != nil {
		fmt.Printf("SQL Error in RegisterVehicleHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to register the vehicle"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with id %d created successfully", id)})
}

func UpdateVehicleInfoHandler(c *gin.Context) {
	var updateVehicleRequest VehicleRequest
	if err := c.BindJSON(&updateVehicleRequest); err != nil {
		fmt.Printf("Error binding POST body in UpdateVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}

	newVehicle := ConvertVehicleRequestToDTO(updateVehicleRequest)
	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE vehicles SET ")
	params := []interface{}{}
	paramsQueryMaps := map[string]interface{}{
		"license_plate": newVehicle.LicensePlate,
		"type":          newVehicle.Type,
		"make":          newVehicle.Make,
		"model":         newVehicle.Model,
		"year":          newVehicle.Year,
		"capacity_kg":   newVehicle.CapacityKg,
		"driver_id":     newVehicle.DriverID,
	}

	for column, param := range paramsQueryMaps {
		if param == "" || param == 0 {
			continue
		}
		queryBuilder.WriteString(fmt.Sprintf("%s = ?, ", column))
		params = append(params, param)
	}
	if len(params) == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "nothing to update"})
		return
	}
	params = append(params, newVehicle.ID)
	db := datastore.GetSQLDataStore(c)
	// removing the last character because it contains a leftover comma
	_, err := db.Exec(queryBuilder.String()[:queryBuilder.Len()-1]+" WHERE id = ?", params...)
	if err != nil {
		fmt.Printf("SQL Error in UpdateVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to update the vehicle"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with id %s updated successfully", newVehicle.ID)})
}

func DeleteVehicleInfoHandler(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no id param found"})
		return
	}
	query := "DELETE FROM vehicles WHERE id = ?"
	db := datastore.GetSQLDataStore(c)
	if _, err := db.Exec(query, ID); err != nil {
		fmt.Printf("SQL Error in DeleteVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal error while deleting"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with id %s deleted successfully", ID)})
}

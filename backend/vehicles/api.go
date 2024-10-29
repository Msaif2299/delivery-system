package vehicles

import (
	"delivery-system/datastore"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetVehicleHandler(c *gin.Context) {
	ID := c.Param("license_plate")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no license_plate param found"})
		return
	}
	var vehicle *VehicleDTO
	vehicle, err := GetVehicleDetails(c, ID)
	if err != nil {
		fmt.Printf("SQL Error in GetVehicleHandler, err: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ConvertVehicleDTOToRequest(*vehicle))
}

func RegisterVehicleHandler(c *gin.Context) {
	var newVehicleRequest VehicleRequest
	if err := c.BindJSON(&newVehicleRequest); err != nil {
		fmt.Printf("Error binding POST body in RegisterVehicleHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}
	if newVehicleRequest.LicensePlate == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "license plate cannot be empty"})
		return
	}

	newVehicle := ConvertVehicleRequestToDTO(newVehicleRequest)
	query := "INSERT INTO vehicles (license_plate, type, make, model, year, capacity_kg, driver_id) VALUES (?,?,?,?,?,?,?)"
	db := datastore.GetSQLDataStore(c)
	_, err := db.Exec(query, newVehicle.LicensePlate, newVehicle.Type, newVehicle.Make, newVehicle.Model, newVehicle.Year, newVehicle.CapacityKg, newVehicle.DriverID)
	if err != nil {
		fmt.Printf("SQL Error in RegisterVehicleHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to register the vehicle"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with license plate %s created successfully", newVehicle.LicensePlate)})
}

func UpdateVehicleInfoHandler(c *gin.Context) {
	var updateVehicleRequest VehicleRequest
	if err := c.BindJSON(&updateVehicleRequest); err != nil {
		fmt.Printf("Error binding POST body in UpdateVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}
	if updateVehicleRequest.LicensePlate == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "license plate cannot be empty"})
		return
	}

	newVehicle := ConvertVehicleRequestToDTO(updateVehicleRequest)
	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE vehicles SET ")
	params := []interface{}{}
	paramsQueryMaps := map[string]interface{}{
		"type":        newVehicle.Type,
		"make":        newVehicle.Make,
		"model":       newVehicle.Model,
		"year":        newVehicle.Year,
		"capacity_kg": newVehicle.CapacityKg,
		"driver_id":   newVehicle.DriverID,
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
	params = append(params, newVehicle.LicensePlate)
	db := datastore.GetSQLDataStore(c)
	// removing the last character because it contains a leftover comma
	_, err := db.Exec(queryBuilder.String()[:queryBuilder.Len()-1]+" WHERE license_plate = ?", params...)
	if err != nil {
		fmt.Printf("SQL Error in UpdateVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to update the vehicle"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with license plate %s updated successfully", newVehicle.LicensePlate)})
}

func DeleteVehicleInfoHandler(c *gin.Context) {
	ID := c.Param("license_plate")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no license_plate param found"})
		return
	}
	query := "DELETE FROM vehicles WHERE license_plate = ?"
	db := datastore.GetSQLDataStore(c)
	if _, err := db.Exec(query, ID); err != nil {
		fmt.Printf("SQL Error in DeleteVehicleInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal error while deleting"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("vehicle with license_plate %s deleted successfully", ID)})
}

func AssignDriverToVehicleHandler(c *gin.Context) {
	var request struct {
		VehicleLicensePlate string `json:"vehicle_license_plate"`
		DriverLicenseNumber string `json:"driver_license_number"`
	}
	if err := c.BindJSON(&request); err != nil {
		fmt.Printf("Error binding JSON in AssignDriverToVehicleHandler, err: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "malformed request"})
		return
	}
	if request.DriverLicenseNumber == "" || request.VehicleLicensePlate == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "driver license number or vehicle license plate cannot be empty"})
		return
	}
	query := `
		UPDATE vehicles
		SET driver_id = (
			SELECT id FROM drivers WHERE license_number = ?
		)
		WHERE license_plate = ? AND EXISTS (
			SELECT 1 FROM drivers WHERE license_number = ?
		);
	`
	db := datastore.GetSQLDataStore(c)
	rowsAffected, err := db.Exec(query, request.DriverLicenseNumber, request.VehicleLicensePlate, request.DriverLicenseNumber)
	if err != nil {
		fmt.Printf("Error encountered in SQL in AssignDriverToVehicleHandler, err: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to assign driver to vehicle"})
		return
	}
	if rowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "vehicle or driver not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "driver assigned to vehicle successfully"})
}

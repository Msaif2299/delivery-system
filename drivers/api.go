package drivers

import (
	"database/sql"
	"delivery-system/datastore"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetDriverHandler(c *gin.Context) {
	ID := c.Param("license_number")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no license number param found"})
		return
	}
	query := "SELECT * FROM drivers WHERE license_number = ?"
	db := datastore.GetSQLDataStore(c)
	var driver DriverDTO
	row := db.QueryRow(query, ID)
	if err := row.Scan(&driver.ID,
		&driver.FullName,
		&driver.LicenseNumber,
		&driver.PrimaryPhoneNumber,
		&driver.PrimaryPhoneCountryCode,
		&driver.SecondaryPhoneNumber,
		&driver.SecondaryPhoneCountryCode,
		&driver.Email,
		&driver.Status,
	); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "driver license number does not exist"})
			return
		}
		fmt.Printf("SQL Error encountered in GetDriverHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal error while searching"})
		return
	}
	c.JSON(http.StatusFound, ConvertDTOToDriverRequest(driver))
}

func RegisterDriverHandler(c *gin.Context) {
	var newDriverRequest DriverRequest
	if err := c.BindJSON(&newDriverRequest); err != nil {
		fmt.Printf("Error binding POST body in RegisterDriverHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}
	if newDriverRequest.LicenseNumber == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "driver license number cannot be zero"})
		return
	}
	newDriver := ConvertDriverRequestToDTO(newDriverRequest)
	query := "INSERT INTO drivers (" +
		"full_name, " +
		"license_number, " +
		"primary_phone_country_code, " +
		"primary_phone_number, " +
		"secondary_phone_country_code, " +
		"secondary_phone_number, " +
		"email, " +
		"status)" +
		"VALUES (?,?,?,?,?,?,?,?)"
	db := datastore.GetSQLDataStore(c)
	_, err := db.Exec(
		query,
		newDriver.FullName,
		newDriver.LicenseNumber,
		newDriver.PrimaryPhoneCountryCode,
		newDriver.PrimaryPhoneNumber,
		newDriver.SecondaryPhoneCountryCode,
		newDriver.SecondaryPhoneNumber,
		newDriver.Email,
		newDriver.Status,
	)
	if err != nil {
		fmt.Printf("Error encountered in SQL in RegisterDriverHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to register the driver"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("driver with id %s created successfully", newDriver.LicenseNumber)})
}

func UpdateDriverInfoHandler(c *gin.Context) {
	var updateDriverRequest DriverRequest
	if err := c.BindJSON(updateDriverRequest); err != nil {
		fmt.Printf("Error binding POST body in RegisterDriverHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json body is malformed"})
		return
	}
	if updateDriverRequest.LicenseNumber == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "driver license number cannot be zero"})
		return
	}
	newDriver := ConvertDriverRequestToDTO(updateDriverRequest)
	var queryBuilder strings.Builder
	queryBuilder.WriteString("UPDATE drivers SET ")
	params := []interface{}{}
	paramsQueryMaps := map[string]interface{}{
		"full_name":                    newDriver.FullName,
		"license_number":               newDriver.LicenseNumber,
		"primary_phone_country_code":   newDriver.PrimaryPhoneCountryCode,
		"primary_phone_number":         newDriver.PrimaryPhoneNumber,
		"secondary_phone_country_code": newDriver.SecondaryPhoneCountryCode,
		"secondary_phone_number":       newDriver.SecondaryPhoneNumber,
		"email":                        newDriver.Email,
		"status":                       newDriver.Status,
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
	params = append(params, newDriver.LicenseNumber)
	db := datastore.GetSQLDataStore(c)
	_, err := db.Exec(queryBuilder.String()[:queryBuilder.Len()-1]+" WHERE license_number = ?", params...)
	if err != nil {
		fmt.Printf("Error encountered in SQL in UpdateDriverInfo: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "unable to update the driver"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("driver with id %s created successfully", newDriver.LicenseNumber)})
}

// TODO: Change to soft delete on a later date
func DeleteDriverInfoHandler(c *gin.Context) {
	ID := c.Param("license_number")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no license_number param found"})
		return
	}
	query := "DELETE FROM table WHERE license_number = ?"
	db := datastore.GetSQLDataStore(c)
	if _, err := db.Exec(query, ID); err != nil {
		fmt.Printf("SQL Error encountered in DeleteDriverInfoHandler: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "internal error while deleting"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("driver with license_number %s deleted successfully", ID)})
}

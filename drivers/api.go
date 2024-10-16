package drivers

import (
	"database/sql"
	"delivery-system/datastore"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDriverHandler(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no id param found"})
		return
	}
	query := "SELECT * FROM drivers WHERE id = ?"
	db := datastore.GetSQLDataStore(c)
	var driver Driver
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
			c.JSON(http.StatusNotFound, gin.H{"message": "driver id does not exist"})
			return
		}
		fmt.Printf("SQL Error encountered in GetDriverHandler: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error while searching"})
		return
	}
	c.JSON(http.StatusFound, driver)
}

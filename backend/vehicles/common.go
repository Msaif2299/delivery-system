package vehicles

import (
	"database/sql"
	"delivery-system/datastore"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetVehicleDetails(c *gin.Context, vehicleID string) (*VehicleDTO, error) {
	query := "SELECT * FROM vehicles WHERE license_plate = ?"
	db := datastore.GetSQLDataStore(c)
	var vehicle VehicleDTO
	err := db.Get(c.Request.Context(), query, fmt.Sprintf("get-vehicle-%s", vehicleID), &vehicle, vehicleID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Unable to find vehicle %s\n", vehicleID)
			return nil, fmt.Errorf("vehicle does not exist")
		}
		fmt.Printf("SQL Error in CheckIfVehicleExists: %s", err.Error())
		return nil, fmt.Errorf("internal error while searching")
	}
	return &vehicle, nil
}

func ValidateLicensePlateNumber(plate string) error {
	if plate == "" {
		return fmt.Errorf("license plate cannot be empty")
	}
	return nil
}

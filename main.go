package main

import (
	"delivery-system/datastore"
	"delivery-system/drivers"
	"fmt"
)

func getAll(db datastore.SQLDataStore) {
	rows, err := db.Query("SELECT * FROM `drivers`")
	if err != nil {
		fmt.Printf("Error encountered, err: %s\n", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var driver drivers.Driver
		if err := rows.Scan(
			&driver.ID,
			&driver.FullName,
			&driver.LicenseNumber,
			&driver.PrimaryPhoneNumber,
			&driver.PrimaryPhoneCountryCode,
			&driver.SecondaryPhoneNumber,
			&driver.SecondaryPhoneCountryCode,
			&driver.Email,
			&driver.Status,
		); err != nil {
			fmt.Printf("Error encountered while reading from sql db, err: %s \n", err.Error())
			return
		}
		fmt.Printf("Found: %+v\n", driver)
	}
}

func main() {
	var sqlDB datastore.SQLDataStore
	var err error
	sqlDB, err = datastore.NewMySQLDataStore()
	if err != nil {
		fmt.Printf("Could not establish connection, err: %s exiting...", err.Error())
		return
	}
	getAll(sqlDB)
}

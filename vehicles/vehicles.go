package vehicles

type VehicleDTO struct {
	ID           int64   `json:"id"`
	LicensePlate string  `json:"license_plate"`
	Type         *string `json:"type,omitempty"` // e.g., truck, van, car, bike
	Make         *string `json:"make,omitempty"`
	Model        *string `json:"model,omitempty"`
	Year         *int    `json:"year,omitempty"`
	CapacityKg   *int    `json:"capacity_kg,omitempty"`
	DriverID     *int64  `json:"driver_id,omitempty"` // driver may not be assigned, new purchases
}

type VehicleRequest struct {
	ID           int64  `json:"id,omitempty"`
	LicensePlate string `json:"license_plate"`
	Type         string `json:"type,omitempty"` // e.g., truck, van, car, bike
	Make         string `json:"make,omitempty"`
	Model        string `json:"model,omitempty"`
	Year         int    `json:"year,omitempty"`
	CapacityKg   int    `json:"capacity_kg,omitempty"`
	DriverID     int64  `json:"driver_id,omitempty"` // driver may not be assigned, new purchases
}

func ConvertVehicleRequestToDTO(req VehicleRequest) VehicleDTO {
	return VehicleDTO{
		ID:           req.ID,
		LicensePlate: req.LicensePlate,
		Type:         &req.Type,
		Make:         &req.Make,
		Model:        &req.Model,
		Year:         &req.Year,
		CapacityKg:   &req.CapacityKg,
		DriverID:     &req.DriverID,
	}
}

func ConvertVehicleDTOToRequest(dto VehicleDTO) (req VehicleRequest) {
	req.ID = dto.ID
	req.LicensePlate = dto.LicensePlate
	if dto.Type != nil {
		req.Type = *dto.Type
	}
	if dto.Make != nil {
		req.Make = *dto.Make
	}
	if dto.Model != nil {
		req.Model = *dto.Model
	}
	if dto.Year != nil {
		req.Year = *dto.Year
	}
	if dto.CapacityKg != nil {
		req.CapacityKg = *dto.CapacityKg
	}
	if dto.DriverID != nil {
		req.DriverID = *dto.DriverID
	}
	return
}

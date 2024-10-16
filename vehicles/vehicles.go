package vehicles

type VehicleDTO struct {
	ID           string `json:"id"`
	LicensePlate string `json:"license_plate,omitempty"`
	Type         string `json:"type,omitempty"` // e.g., truck, van, car, bike
	Make         string `json:"make,omitempty"`
	Model        string `json:"model,omitempty"`
	Year         int    `json:"year,omitempty"`
	CapacityKg   int    `json:"capacity_kg,omitempty"`
	DriverID     int64  `json:"driver_id,omitempty"` // driver may not be assigned, new purchases
}

type VehicleRequest struct {
	ID           string `json:"id"`
	LicensePlate string `json:"license_plate,omitempty"`
	Type         string `json:"type,omitempty"` // e.g., truck, van, car, bike
	Make         string `json:"make,omitempty"`
	Model        string `json:"model,omitempty"`
	Year         int    `json:"year,omitempty"`
	CapacityKg   int    `json:"capacity_kg,omitempty"`
	DriverID     int64  `json:"driver_id,omitempty"` // driver may not be assigned, new purchases
}

func ConvertVehicleRequestToDTO(req VehicleRequest) VehicleDTO {
	return VehicleDTO(req)
}

func ConvertVehicleDTOToRequest(dto VehicleDTO) VehicleRequest {
	return VehicleRequest(dto)
}

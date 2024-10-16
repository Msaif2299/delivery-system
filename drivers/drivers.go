package drivers

// DriverDTO is used to transfer data into and out of SQL
type DriverDTO struct {
	ID                        int64  `json:"id"`
	FullName                  string `json:"full_name,omitempty"`
	LicenseNumber             string `json:"license_number,omitempty"`
	PrimaryPhoneNumber        string `json:"primary_phone_number,omitempty"`
	PrimaryPhoneCountryCode   string `json:"primary_phone_country_code,omitempty"`
	SecondaryPhoneNumber      string `json:"secondary_phone_number,omitempty"`
	SecondaryPhoneCountryCode string `json:"secondary_phone_country_code,omitempty"`
	Email                     string `json:"email,omitempty"`
	Status                    uint8  `json:"status,omitempty"`
}

// DriverRequest is used to get the data from POST and PUT requests
type DriverRequest struct {
	ID                        int64  `json:"id"`
	FullName                  string `json:"full_name,omitempty"`
	LicenseNumber             string `json:"license_number,omitempty"`
	PrimaryPhoneNumber        string `json:"primary_phone_number,omitempty"`
	PrimaryPhoneCountryCode   string `json:"primary_phone_country_code,omitempty"`
	SecondaryPhoneNumber      string `json:"secondary_phone_number,omitempty"`
	SecondaryPhoneCountryCode string `json:"secondary_phone_country_code,omitempty"`
	Email                     string `json:"email,omitempty"`
	Status                    uint8  `json:"status,omitempty"`
}

func ConvertDriverRequestToDTO(req DriverRequest) DriverDTO {
	return DriverDTO(req)
}

func ConvertDTOToDriverRequest(dto DriverDTO) DriverRequest {
	return DriverRequest(dto)
}

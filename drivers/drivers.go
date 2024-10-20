package drivers

// DriverDTO is used to transfer data into and out of SQL
type DriverDTO struct {
	ID                        int64   `json:"id"`
	FullName                  *string `json:"full_name,omitempty"`
	LicenseNumber             string  `json:"license_number"`
	PrimaryPhoneNumber        *string `json:"primary_phone_number,omitempty"`
	PrimaryPhoneCountryCode   *string `json:"primary_phone_country_code,omitempty"`
	SecondaryPhoneNumber      *string `json:"secondary_phone_number,omitempty"`
	SecondaryPhoneCountryCode *string `json:"secondary_phone_country_code,omitempty"`
	Email                     *string `json:"email,omitempty"`
	Status                    *uint8  `json:"status,omitempty"`
}

// DriverRequest is used to get the data from POST and PUT requests
type DriverRequest struct {
	ID                        int64  `json:"id,omitempty"`
	FullName                  string `json:"full_name,omitempty"`
	LicenseNumber             string `json:"license_number"`
	PrimaryPhoneNumber        string `json:"primary_phone_number,omitempty"`
	PrimaryPhoneCountryCode   string `json:"primary_phone_country_code,omitempty"`
	SecondaryPhoneNumber      string `json:"secondary_phone_number,omitempty"`
	SecondaryPhoneCountryCode string `json:"secondary_phone_country_code,omitempty"`
	Email                     string `json:"email,omitempty"`
	Status                    uint8  `json:"status,omitempty"`
}

func ConvertDriverRequestToDTO(req DriverRequest) DriverDTO {
	return DriverDTO{
		ID:                        req.ID,
		FullName:                  &req.FullName,
		LicenseNumber:             req.LicenseNumber,
		PrimaryPhoneNumber:        &req.PrimaryPhoneNumber,
		PrimaryPhoneCountryCode:   &req.PrimaryPhoneCountryCode,
		SecondaryPhoneNumber:      &req.SecondaryPhoneNumber,
		SecondaryPhoneCountryCode: &req.SecondaryPhoneCountryCode,
		Email:                     &req.Email,
		Status:                    &req.Status,
	}
}

func ConvertDTOToDriverRequest(dto DriverDTO) DriverRequest {
	req := DriverRequest{
		ID:            dto.ID,
		LicenseNumber: dto.LicenseNumber,
		Status:        *dto.Status,
	}
	if dto.FullName != nil {
		req.FullName = *dto.FullName
	}
	if dto.PrimaryPhoneNumber != nil {
		req.PrimaryPhoneNumber = *dto.PrimaryPhoneNumber
	}
	if dto.PrimaryPhoneCountryCode != nil {
		req.PrimaryPhoneCountryCode = *dto.PrimaryPhoneCountryCode
	}
	if dto.SecondaryPhoneNumber != nil {
		req.SecondaryPhoneNumber = *dto.SecondaryPhoneNumber
	}
	if dto.SecondaryPhoneCountryCode != nil {
		req.SecondaryPhoneCountryCode = *dto.SecondaryPhoneCountryCode
	}
	if dto.Email != nil {
		req.Email = *dto.Email
	}
	if dto.Status != nil {
		req.Status = *dto.Status
	}
	return req
}

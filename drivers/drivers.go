package drivers

type Driver struct {
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

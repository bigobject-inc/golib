package license

// LicenseData License data
type LicenseData struct {
	AuthID         string `json:"authID"`
	ClientID       string `json:"clientID"`
	ContactEmail   string `json:"contactEmail"`
	ContactName    string `json:"contactName"`
	ContactPhone   string `json:"contactPhone"`
	Desc           string `json:"desc"`
	ExpiredAt      string `json:"expiredAt"`
	Level          string `json:"level"`
	Organization   string `json:"organization"`
	OrganizationID string `json:"organizationID"`
}

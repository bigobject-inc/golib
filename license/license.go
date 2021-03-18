package license

// License BigObject License
type License interface {
	Decrypt(clientID string, LicenseData []byte) (LicenseData, error)
	Encrypt(l LicenseData) ([]byte, error)
	NewClientID(organization, organizationID string) string
	ValidateClientID(clientID, organization, organizationID string) bool
}

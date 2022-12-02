package elements

// The OpenAsAdminOrSystemService element is for internal use only. This element is not used by clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/openasadminorsystemservice
type OpenAsAdminOrSystemService struct {
	// Not intended for client use.
	LogonType *string `xml:"LogonType,attr"`
}

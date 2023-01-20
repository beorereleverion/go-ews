package elements

// The OpenAsAdminOrSystemService element is for internal use only. This element is not used by clients.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/openasadminorsystemservice
import "encoding/xml"

type OpenAsAdminOrSystemService struct {
	XMLName xml.Name

	// Not intended for client use.
	LogonType *string `xml:"LogonType,attr"`
}

func (O *OpenAsAdminOrSystemService) SetForMarshal() {
	O.XMLName.Local = "t:OpenAsAdminOrSystemService"
}

func (O *OpenAsAdminOrSystemService) GetSchema() *Schema {
	return &SchemaTypes
}

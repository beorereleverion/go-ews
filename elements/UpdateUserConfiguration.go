package elements

// The UpdateUserConfiguration element represents a request to update a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateuserconfiguration
import "encoding/xml"

type UpdateUserConfiguration struct {
	XMLName xml.Name

	// The UserConfiguration element defines a single user configuration object.
	UserConfiguration *UserConfiguration `xml:"UserConfiguration"`
}

func (U *UpdateUserConfiguration) SetForMarshal() {
	U.XMLName.Local = "m:UpdateUserConfiguration"
}

func (U *UpdateUserConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}

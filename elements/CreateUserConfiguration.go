package elements

// The CreateUserConfiguration element represents a request to create a user configuration object.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createuserconfiguration
import "encoding/xml"

type CreateUserConfiguration struct {
	XMLName xml.Name

	// The UserConfiguration element defines a single user configuration object.
	UserConfiguration *UserConfiguration `xml:"UserConfiguration"`
}

func (C *CreateUserConfiguration) SetForMarshal() {
	C.XMLName.Local = "m:CreateUserConfiguration"
}

func (C *CreateUserConfiguration) GetSchema() *Schema {
	return &SchemaMessages
}

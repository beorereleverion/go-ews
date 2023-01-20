package elements

// The Email element identifies the email address of the user whose photo is requested in the GetUserPhoto operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/email-string
import "encoding/xml"

type EmailString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EmailString) SetForMarshal() {
	E.XMLName.Local = "t:Email"
}

func (E *EmailString) GetSchema() *Schema {
	return &SchemaTypes
}

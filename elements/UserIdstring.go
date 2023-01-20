package elements

// The UserId element specifies the user identifier of an email user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/userid-string
import "encoding/xml"

type UserIdstring struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UserIdstring) SetForMarshal() {
	U.XMLName.Local = "t:UserId"
}

func (U *UserIdstring) GetSchema() *Schema {
	return &SchemaTypes
}

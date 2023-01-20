package elements

// The EmailUser element specifies an email recipient.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/emailuser
import "encoding/xml"

type EmailUser struct {
	XMLName xml.Name

	// The Name element specifies a search refiner name.
	Name *Namestring `xml:"Name"`
	// The UserId element specifies the user identifier of an email user.
	UserId *UserIdstring `xml:"UserId"`
}

func (E *EmailUser) SetForMarshal() {
	E.XMLName.Local = "t:EmailUser"
}

func (E *EmailUser) GetSchema() *Schema {
	return &SchemaTypes
}

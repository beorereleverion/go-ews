package elements

// The DistinguishedUser element identifies Anonymous and Default user accounts for delegate access. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/distinguisheduser
import "encoding/xml"

type DistinguishedUser struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// anonymous
	DistinguishedUseranonymous string = `anonymous`
	// default
	DistinguishedUserdefault string = `default`
)

func (D *DistinguishedUser) SetForMarshal() {
	D.XMLName.Local = "t:DistinguishedUser"
}

func (D *DistinguishedUser) GetSchema() *Schema {
	return &SchemaTypes
}

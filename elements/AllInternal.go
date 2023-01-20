package elements

// The AllInternal element evaluates to true if all recipients of an e-mail message are internal to the sender's organization.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/allinternal
import "encoding/xml"

type AllInternal struct {
	XMLName xml.Name
	TEXT    struct{} `xml:",chardata"`
}

func (A *AllInternal) SetForMarshal() {
	A.XMLName.Local = "t:AllInternal"
}

func (A *AllInternal) GetSchema() *Schema {
	return &SchemaTypes
}

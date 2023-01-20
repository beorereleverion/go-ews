package elements

// The AddressEntity element specifies a single address entity.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addressentity
import "encoding/xml"

type AddressEntity struct {
	XMLName xml.Name

	// The Address element represents the e-mail address of the mailbox user.
	Address *Addressstring `xml:"Address"`
	// The Position element specifies the position of an entity extracted from a message.
	Position *Position `xml:"Position"`
}

func (A *AddressEntity) SetForMarshal() {
	A.XMLName.Local = "t:AddressEntity"
}

func (A *AddressEntity) GetSchema() *Schema {
	return &SchemaTypes
}

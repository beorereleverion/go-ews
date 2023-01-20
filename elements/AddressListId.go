package elements

// The AddressListId element specifies the identifier of an address list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addresslistid
import "encoding/xml"

type AddressListId struct {
	XMLName xml.Name

	// A string address list identifier. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (A *AddressListId) SetForMarshal() {
	A.XMLName.Local = "t:AddressListId"
}

func (A *AddressListId) GetSchema() *Schema {
	return &SchemaTypes
}

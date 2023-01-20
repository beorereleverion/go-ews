package elements

// The MemberCorrelationKey element specifies the identifiers of the contacts that are part of the instant messaging (IM) group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/membercorrelationkey
import "encoding/xml"

type MemberCorrelationKey struct {
	XMLName xml.Name

	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"ItemId"`
}

func (M *MemberCorrelationKey) SetForMarshal() {
	M.XMLName.Local = "t:MemberCorrelationKey"
}

func (M *MemberCorrelationKey) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The ExchangeStoreId element specifies the instant messaging (IM) group identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/exchangestoreid
import "encoding/xml"

type ExchangeStoreId struct {
	XMLName xml.Name

	// The text value of the ChangeKey attribute is the change key of the group.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// The text value of the Id attribute is the identifier of the group.
	Id *string `xml:"Id,attr"`
}

func (E *ExchangeStoreId) SetForMarshal() {
	E.XMLName.Local = "t:ExchangeStoreId"
}

func (E *ExchangeStoreId) GetSchema() *Schema {
	return &SchemaTypes
}

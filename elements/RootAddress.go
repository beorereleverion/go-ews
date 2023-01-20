package elements

// The RootAddress element represents the first address that starts the event for a RecipientTrackingEvent event.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/rootaddress
import "encoding/xml"

type RootAddress struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (R *RootAddress) SetForMarshal() {
	R.XMLName.Local = "t:RootAddress"
}

func (R *RootAddress) GetSchema() *Schema {
	return &SchemaTypes
}

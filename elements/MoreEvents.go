package elements

// The MoreEvents element indicates whether there are more events in the queue to be delivered to the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moreevents
import "encoding/xml"

type MoreEvents struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (M *MoreEvents) SetForMarshal() {
	M.XMLName.Local = "t:MoreEvents"
}

func (M *MoreEvents) GetSchema() *Schema {
	return &SchemaTypes
}

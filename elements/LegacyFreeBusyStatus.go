package elements

// The LegacyFreeBusyStatus element represents the free/busy status of the calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/legacyfreebusystatus
import "encoding/xml"

type LegacyFreeBusyStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (L *LegacyFreeBusyStatus) SetForMarshal() {
	L.XMLName.Local = "t:LegacyFreeBusyStatus"
}

func (L *LegacyFreeBusyStatus) GetSchema() *Schema {
	return &SchemaTypes
}

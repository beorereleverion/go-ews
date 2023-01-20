package elements

// The GetServerTimeZones element is the root element in a request to retrieve time zone definitions from the Exchange server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getservertimezones
import "encoding/xml"

type GetServerTimeZones struct {
	XMLName xml.Name

	// The Ids element contains an array of time zone definition identifiers.
	Ids *Ids `xml:"Ids"`
	// Specifies whether the GetServerTimeZones operation returns the complete definition or only the name and identifier for each time zone. This attribute is optional. The default value is true.
	ReturnFullTimeZoneData *string `xml:"ReturnFullTimeZoneData,attr"`
}

func (G *GetServerTimeZones) SetForMarshal() {
	G.XMLName.Local = "m:GetServerTimeZones"
}

func (G *GetServerTimeZones) GetSchema() *Schema {
	return &SchemaMessages
}

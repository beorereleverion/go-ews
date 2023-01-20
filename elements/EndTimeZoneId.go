package elements

// The EndTimeZoneId element specifies the time zone in which a meeting takes place.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtimezoneid
import "encoding/xml"

type EndTimeZoneId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EndTimeZoneId) SetForMarshal() {
	E.XMLName.Local = "t:EndTimeZoneId"
}

func (E *EndTimeZoneId) GetSchema() *Schema {
	return &SchemaTypes
}

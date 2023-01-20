package elements

// The StartTimeZoneId element specifies the time zone in which a meeting takes place.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttimezoneid
import "encoding/xml"

type StartTimeZoneId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *StartTimeZoneId) SetForMarshal() {
	S.XMLName.Local = "t:StartTimeZoneId"
}

func (S *StartTimeZoneId) GetSchema() *Schema {
	return &SchemaTypes
}

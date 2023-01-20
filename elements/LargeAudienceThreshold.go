package elements

// The LargeAudienceThreshold element represents the large audience threshold for a client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/largeaudiencethreshold
import "encoding/xml"

type LargeAudienceThreshold struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (L *LargeAudienceThreshold) SetForMarshal() {
	L.XMLName.Local = "t:LargeAudienceThreshold"
}

func (L *LargeAudienceThreshold) GetSchema() *Schema {
	return &SchemaTypes
}

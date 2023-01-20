package elements

// The GoodThreshold element specifies the percentage of attendees that must have the time period open in order for the time period to qualify as a Good suggested meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/goodthreshold
import "encoding/xml"

type GoodThreshold struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (G *GoodThreshold) SetForMarshal() {
	G.XMLName.Local = "t:GoodThreshold"
}

func (G *GoodThreshold) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The PreviousWatermark element represents the watermark of the latest event that was successfully communicated to the client for the subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/previouswatermark
import "encoding/xml"

type PreviousWatermark struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PreviousWatermark) SetForMarshal() {
	P.XMLName.Local = "t:PreviousWatermark"
}

func (P *PreviousWatermark) GetSchema() *Schema {
	return &SchemaTypes
}

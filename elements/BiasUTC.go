package elements

// The Bias element represents the general offset from Coordinated Universal Time (UTC). This value is in minutes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bias-utc
import "encoding/xml"

type BiasUTC struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (B *BiasUTC) SetForMarshal() {
	B.XMLName.Local = "t:Bias"
}

func (B *BiasUTC) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The Bias element represents the offset from the Coordinated Universal Time (UTC) offset identified by the Bias (UTC) element for standard time and daylight saving time. This value is in minutes.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bias
import "encoding/xml"

type Bias struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (B *Bias) SetForMarshal() {
	B.XMLName.Local = "t:Bias"
}

func (B *Bias) GetSchema() *Schema {
	return &SchemaTypes
}

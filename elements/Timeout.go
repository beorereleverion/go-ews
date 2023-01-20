package elements

// The Timeout element represents the duration, in minutes, that the subscription can remain idle without a GetEvents request from the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timeout
import "encoding/xml"

type Timeout struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *Timeout) SetForMarshal() {
	T.XMLName.Local = "t:Timeout"
}

func (T *Timeout) GetSchema() *Schema {
	return &SchemaTypes
}

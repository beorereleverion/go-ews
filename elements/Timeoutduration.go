package elements

// The Timeout element specifies the length of time before a pull subscription is timed out by the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/timeout-duration
import "encoding/xml"

type Timeoutduration struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *Timeoutduration) SetForMarshal() {
	T.XMLName.Local = "t:Timeout"
}

func (T *Timeoutduration) GetSchema() *Schema {
	return &SchemaTypes
}

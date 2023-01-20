package elements

// The StatusFrequency element represents the maximum timeout value, in minutes, in which retries are attempted by the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/statusfrequency
import "encoding/xml"

type StatusFrequency struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (S *StatusFrequency) SetForMarshal() {
	S.XMLName.Local = "t:StatusFrequency"
}

func (S *StatusFrequency) GetSchema() *Schema {
	return &SchemaTypes
}

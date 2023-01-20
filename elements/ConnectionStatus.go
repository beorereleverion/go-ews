package elements

// The ConnectionStatus element provides a text description of the status of a streaming subscription.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/connectionstatus
import "encoding/xml"

type ConnectionStatus struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *ConnectionStatus) SetForMarshal() {
	C.XMLName.Local = "m:ConnectionStatus"
}

func (C *ConnectionStatus) GetSchema() *Schema {
	return &SchemaMessages
}

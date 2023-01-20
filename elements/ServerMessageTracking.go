package elements

// The Server element represents the physical server where the event occurred.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/server-messagetracking
import "encoding/xml"

type ServerMessageTracking struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *ServerMessageTracking) SetForMarshal() {
	S.XMLName.Local = "t:Server"
}

func (S *ServerMessageTracking) GetSchema() *Schema {
	return &SchemaTypes
}

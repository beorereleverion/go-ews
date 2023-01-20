package elements

// The PreviousHopServer element represents the previous server name that accepted the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/previoushopserver
import "encoding/xml"

type PreviousHopServer struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PreviousHopServer) SetForMarshal() {
	P.XMLName.Local = "t:PreviousHopServer"
}

func (P *PreviousHopServer) GetSchema() *Schema {
	return &SchemaTypes
}

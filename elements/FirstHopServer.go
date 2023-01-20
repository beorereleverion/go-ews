package elements

// The FirstHopServer element contains the name of the server in the forest that first accepted the message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firsthopserver
import "encoding/xml"

type FirstHopServer struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FirstHopServer) SetForMarshal() {
	F.XMLName.Local = "t:FirstHopServer"
}

func (F *FirstHopServer) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The GetRoomLists element is the root element in a request for a list of email addresses that represent a list of available rooms.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getroomlists
import "encoding/xml"

type GetRoomLists struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (G *GetRoomLists) SetForMarshal() {
	G.XMLName.Local = "m:GetRoomLists"
}

func (G *GetRoomLists) GetSchema() *Schema {
	return &SchemaMessages
}

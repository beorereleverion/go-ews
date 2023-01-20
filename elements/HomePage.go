package elements

// The HomePage element specifies the URL that will be the default home page for the managed folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/homepage
import "encoding/xml"

type HomePage struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (H *HomePage) SetForMarshal() {
	H.XMLName.Local = "t:HomePage"
}

func (H *HomePage) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The ActionUrl element identifies the URL that the user should navigate to, in order to fix an issue indicated by the AppStatus element.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/actionurl
import "encoding/xml"

type ActionUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (A *ActionUrl) SetForMarshal() {
	A.XMLName.Local = "t:ActionUrl"
}

func (A *ActionUrl) GetSchema() *Schema {
	return &SchemaTypes
}

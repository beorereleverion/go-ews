package elements

// The EndNodeUrl element specifies the URL for the mail app in the Office Store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endnodeurl
import "encoding/xml"

type EndNodeUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (E *EndNodeUrl) SetForMarshal() {
	E.XMLName.Local = "t:EndNodeUrl"
}

func (E *EndNodeUrl) GetSchema() *Schema {
	return &SchemaTypes
}

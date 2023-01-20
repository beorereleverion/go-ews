package elements

// The OwaLink element specifies the link to preview an item in Microsoft Outlook Web App.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/owalink
import "encoding/xml"

type OwaLink struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (O *OwaLink) SetForMarshal() {
	O.XMLName.Local = "t:OwaLink"
}

func (O *OwaLink) GetSchema() *Schema {
	return &SchemaTypes
}

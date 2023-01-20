package elements

// The WebClientEditFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to edit an item in Outlook Web App.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/webclienteditformquerystring
import "encoding/xml"

type WebClientEditFormQueryString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *WebClientEditFormQueryString) SetForMarshal() {
	W.XMLName.Local = "t:WebClientEditFormQueryString"
}

func (W *WebClientEditFormQueryString) GetSchema() *Schema {
	return &SchemaTypes
}

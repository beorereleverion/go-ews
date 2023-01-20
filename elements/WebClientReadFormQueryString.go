package elements

// The WebClientReadFormQueryString element represents a URL to concatenate to the Outlook Web App endpoint to read an item in Outlook Web App.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/webclientreadformquerystring
import "encoding/xml"

type WebClientReadFormQueryString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (W *WebClientReadFormQueryString) SetForMarshal() {
	W.XMLName.Local = "t:WebClientReadFormQueryString"
}

func (W *WebClientReadFormQueryString) GetSchema() *Schema {
	return &SchemaTypes
}

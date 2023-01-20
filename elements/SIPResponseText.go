package elements

// The SIPResponseText element specifies the SIP response text.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sipresponsetext
import "encoding/xml"

type SIPResponseText struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *SIPResponseText) SetForMarshal() {
	S.XMLName.Local = "t:SIPResponseText"
}

func (S *SIPResponseText) GetSchema() *Schema {
	return &SchemaTypes
}

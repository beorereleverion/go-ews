package elements

// The YomiCompanyName element specifies the phonetic Japanese company name of the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/yomicompanyname
import "encoding/xml"

type YomiCompanyName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (Y *YomiCompanyName) SetForMarshal() {
	Y.XMLName.Local = "t:YomiCompanyName"
}

func (Y *YomiCompanyName) GetSchema() *Schema {
	return &SchemaTypes
}

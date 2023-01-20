package elements

// The WeddingAnniversaries element specifies an array of wedding anniversary dates, stored as strings, and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weddinganniversaries
import "encoding/xml"

type WeddingAnniversaries struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (W *WeddingAnniversaries) SetForMarshal() {
	W.XMLName.Local = "t:WeddingAnniversaries"
}

func (W *WeddingAnniversaries) GetSchema() *Schema {
	return &SchemaTypes
}

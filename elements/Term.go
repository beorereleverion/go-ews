package elements

// The Term element specifies a highlighted term in a FindConversation or FindItem response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/term
import "encoding/xml"

type Term struct {
	XMLName xml.Name

	// The Scope element specifies the string to be highlighted.
	Scope *ScopeHighlightTermType `xml:"Scope"`
	// The Value element contains the value of an extended property.
	Value *Value `xml:"Value"`
}

func (T *Term) SetForMarshal() {
	T.XMLName.Local = "t:Term"
}

func (T *Term) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The Scope element specifies the string to be highlighted.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/scope-highlighttermtype
import "encoding/xml"

type ScopeHighlightTermType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *ScopeHighlightTermType) SetForMarshal() {
	S.XMLName.Local = "t:Scope"
}

func (S *ScopeHighlightTermType) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The AssistantNames element specifies an array of assistant names and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assistantnames
import "encoding/xml"

type AssistantNames struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (A *AssistantNames) SetForMarshal() {
	A.XMLName.Local = "t:AssistantNames"
}

func (A *AssistantNames) GetSchema() *Schema {
	return &SchemaTypes
}

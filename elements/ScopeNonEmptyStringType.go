package elements

// The Scope element specifies the scope of the message tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/scope-nonemptystringtype
import "encoding/xml"

type ScopeNonEmptyStringType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

const (
	// The message tracking scopes spans across a forest.
	ScopeNonEmptyStringTypeForest string = `Forest`
	// The message tracking scopes spans across an organization.
	ScopeNonEmptyStringTypeOrganization string = `Organization`
	// The message tracking scopes spans across a site.
	ScopeNonEmptyStringTypeSite string = `Site`
)

func (S *ScopeNonEmptyStringType) SetForMarshal() {
	S.XMLName.Local = "m:Scope"
}

func (S *ScopeNonEmptyStringType) GetSchema() *Schema {
	return &SchemaMessages
}

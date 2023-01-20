package elements

// The Scope (ClientAccessTokenRequestType) element specifies a token scope.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/scope-clientaccesstokenrequesttype
import "encoding/xml"

type ScopeClientAccessTokenRequestType struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *ScopeClientAccessTokenRequestType) SetForMarshal() {
	S.XMLName.Local = "t:Scope"
}

func (S *ScopeClientAccessTokenRequestType) GetSchema() *Schema {
	return &SchemaTypes
}

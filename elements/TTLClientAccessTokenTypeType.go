package elements

// The TTL element indicates the time to live value for the token.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ttl-clientaccesstokentypetype
import "encoding/xml"

type TTLClientAccessTokenTypeType struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TTLClientAccessTokenTypeType) SetForMarshal() {
	T.XMLName.Local = "t:TTL"
}

func (T *TTLClientAccessTokenTypeType) GetSchema() *Schema {
	return &SchemaTypes
}

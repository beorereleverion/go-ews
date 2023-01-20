package elements

// The TTL element specifies the time, in minutes, that the token is valid.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ttl
import "encoding/xml"

type TTL struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TTL) SetForMarshal() {
	T.XMLName.Local = "t:TTL"
}

func (T *TTL) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The ServerHint element represents the starting point for tracking a message in a remote site or forest.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/serverhint
import "encoding/xml"

type ServerHint struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (S *ServerHint) SetForMarshal() {
	S.XMLName.Local = "m:ServerHint"
}

func (S *ServerHint) GetSchema() *Schema {
	return &SchemaMessages
}

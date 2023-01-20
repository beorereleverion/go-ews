package elements

// The UniqueHash element specifies a unique hash value used for de-duplication.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquehash
import "encoding/xml"

type UniqueHash struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UniqueHash) SetForMarshal() {
	U.XMLName.Local = "t:UniqueHash"
}

func (U *UniqueHash) GetSchema() *Schema {
	return &SchemaTypes
}

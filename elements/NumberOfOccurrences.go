package elements

// The NumberOfOccurrences element contains the number of occurrences of a recurring item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/numberofoccurrences
import "encoding/xml"

type NumberOfOccurrences struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (N *NumberOfOccurrences) SetForMarshal() {
	N.XMLName.Local = "t:NumberOfOccurrences"
}

func (N *NumberOfOccurrences) GetSchema() *Schema {
	return &SchemaTypes
}

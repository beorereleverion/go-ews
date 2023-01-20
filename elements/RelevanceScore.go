package elements

// The RelevanceScore element specifies an integer that represents how relevant the associated persona is to the client.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relevancescore
import "encoding/xml"

type RelevanceScore struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (R *RelevanceScore) SetForMarshal() {
	R.XMLName.Local = "t:RelevanceScore"
}

func (R *RelevanceScore) GetSchema() *Schema {
	return &SchemaTypes
}

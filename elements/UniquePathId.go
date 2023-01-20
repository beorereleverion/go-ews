package elements

// The UniquePathId element represents a string that is different for each path in a tracking report.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/uniquepathid
import "encoding/xml"

type UniquePathId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (U *UniquePathId) SetForMarshal() {
	U.XMLName.Local = "t:UniquePathId"
}

func (U *UniquePathId) GetSchema() *Schema {
	return &SchemaTypes
}

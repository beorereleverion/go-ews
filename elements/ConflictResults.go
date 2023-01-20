package elements

// The ConflictResults element contains the number of conflicts in an UpdateItem operation response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/conflictresults
import "encoding/xml"

type ConflictResults struct {
	XMLName xml.Name

	// The Count element contains the number of conflicts in an UpdateItem operation response.
	Count *Count `xml:"Count"`
}

func (C *ConflictResults) SetForMarshal() {
	C.XMLName.Local = "t:ConflictResults"
}

func (C *ConflictResults) GetSchema() *Schema {
	return &SchemaTypes
}

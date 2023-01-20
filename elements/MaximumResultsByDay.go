package elements

// The MaximumResultsByDay element specifies the number of suggested meeting times per a day returned in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maximumresultsbyday
import "encoding/xml"

type MaximumResultsByDay struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaximumResultsByDay) SetForMarshal() {
	M.XMLName.Local = "t:MaximumResultsByDay"
}

func (M *MaximumResultsByDay) GetSchema() *Schema {
	return &SchemaTypes
}

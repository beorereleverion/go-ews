package elements

// The MaximumNonWorkHourResultsByDay element specifies the number of suggested results for meeting times outside regular working hours per day.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/maximumnonworkhourresultsbyday
import "encoding/xml"

type MaximumNonWorkHourResultsByDay struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MaximumNonWorkHourResultsByDay) SetForMarshal() {
	M.XMLName.Local = "t:MaximumNonWorkHourResultsByDay"
}

func (M *MaximumNonWorkHourResultsByDay) GetSchema() *Schema {
	return &SchemaTypes
}

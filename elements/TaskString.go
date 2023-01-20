package elements

// The TaskString element contains a suggested task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/taskstring
import "encoding/xml"

type TaskString struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (T *TaskString) SetForMarshal() {
	T.XMLName.Local = "t:TaskString"
}

func (T *TaskString) GetSchema() *Schema {
	return &SchemaTypes
}

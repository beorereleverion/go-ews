package elements

// The TaskTimedOut element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/tasktimedout
import "encoding/xml"

type TaskTimedOut struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (T *TaskTimedOut) SetForMarshal() {
	T.XMLName.Local = "t:TaskTimedOut"
}

func (T *TaskTimedOut) GetSchema() *Schema {
	return &SchemaTypes
}

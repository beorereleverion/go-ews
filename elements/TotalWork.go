package elements

// The TotalWork element contains a description of how much work is associated with a task.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/totalwork
import "encoding/xml"

type TotalWork struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (T *TotalWork) SetForMarshal() {
	T.XMLName.Local = "t:TotalWork"
}

func (T *TotalWork) GetSchema() *Schema {
	return &SchemaTypes
}

package elements

// The OperationIndex element specifies the index of the operation in the request that caused the rule operation error.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/operationindex
import "encoding/xml"

type OperationIndex struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (O *OperationIndex) SetForMarshal() {
	O.XMLName.Local = "m:OperationIndex"
}

func (O *OperationIndex) GetSchema() *Schema {
	return &SchemaMessages
}

package elements

// The RecognitionId element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recognitionid
import "encoding/xml"

type RecognitionId struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (R *RecognitionId) SetForMarshal() {
	R.XMLName.Local = "t:RecognitionId"
}

func (R *RecognitionId) GetSchema() *Schema {
	return &SchemaTypes
}

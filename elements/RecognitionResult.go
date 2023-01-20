package elements

// The RecognitionResult element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recognitionresult
import "encoding/xml"

type RecognitionResult struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (R *RecognitionResult) SetForMarshal() {
	R.XMLName.Local = "t:RecognitionResult"
}

func (R *RecognitionResult) GetSchema() *Schema {
	return &SchemaTypes
}

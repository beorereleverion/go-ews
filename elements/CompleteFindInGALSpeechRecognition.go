package elements

// The CompleteFindInGALSpeechRecognition element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/completefindingalspeechrecognition
import "encoding/xml"

type CompleteFindInGALSpeechRecognition struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (C *CompleteFindInGALSpeechRecognition) SetForMarshal() {
	C.XMLName.Local = "m:CompleteFindInGALSpeechRecognition"
}

func (C *CompleteFindInGALSpeechRecognition) GetSchema() *Schema {
	return &SchemaMessages
}

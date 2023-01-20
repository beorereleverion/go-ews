package elements

// The CompleteFindInGALSpeechRecognitionResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/completefindingalspeechrecognitionresponse
import "encoding/xml"

type CompleteFindInGALSpeechRecognitionResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (C *CompleteFindInGALSpeechRecognitionResponse) SetForMarshal() {
	C.XMLName.Local = "m:CompleteFindInGALSpeechRecognitionResponse"
}

func (C *CompleteFindInGALSpeechRecognitionResponse) GetSchema() *Schema {
	return &SchemaMessages
}

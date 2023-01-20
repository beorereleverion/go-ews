package elements

// The StartFindInGALSpeechRecognition element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startfindingalspeechrecognition
import "encoding/xml"

type StartFindInGALSpeechRecognition struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (S *StartFindInGALSpeechRecognition) SetForMarshal() {
	S.XMLName.Local = "m:StartFindInGALSpeechRecognition"
}

func (S *StartFindInGALSpeechRecognition) GetSchema() *Schema {
	return &SchemaMessages
}

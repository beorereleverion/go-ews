package elements

// The StartFindInGALSpeechRecognitionResponse element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startfindingalspeechrecognitionresponse
import "encoding/xml"

type StartFindInGALSpeechRecognitionResponse struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (S *StartFindInGALSpeechRecognitionResponse) SetForMarshal() {
	S.XMLName.Local = "m:StartFindInGALSpeechRecognitionResponse"
}

func (S *StartFindInGALSpeechRecognitionResponse) GetSchema() *Schema {
	return &SchemaMessages
}

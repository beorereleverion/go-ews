package elements

// The PlayOnPhoneDialString element identifies the Play-on-Phone dial string.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/playonphonedialstring-exchange-web-services
import "encoding/xml"

type PlayOnPhoneDialStringExchangeWebServices struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (P *PlayOnPhoneDialStringExchangeWebServices) SetForMarshal() {
	P.XMLName.Local = "t:PlayOnPhoneDialString"
}

func (P *PlayOnPhoneDialStringExchangeWebServices) GetSchema() *Schema {
	return &SchemaTypes
}

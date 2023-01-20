package elements

// The ProcessRightAway element indicates whether the response is sent as soon as the action starts processing on the server or whether the response is sent after the action has completed. This element must be present for the response to be sent asynchronous to the requested action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/processrightaway
import "encoding/xml"

type ProcessRightAway struct {
	XMLName xml.Name
	TEXT    bool `xml:",chardata"`
}

func (P *ProcessRightAway) SetForMarshal() {
	P.XMLName.Local = "t:ProcessRightAway"
}

func (P *ProcessRightAway) GetSchema() *Schema {
	return &SchemaTypes
}

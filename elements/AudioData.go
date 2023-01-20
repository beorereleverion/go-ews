package elements

// The AudioData element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/audiodata
import "encoding/xml"

type AudioData struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

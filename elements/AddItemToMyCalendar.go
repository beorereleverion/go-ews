package elements

// The AddItemToMyCalendar element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/additemtomycalendar
import "encoding/xml"

type AddItemToMyCalendar struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

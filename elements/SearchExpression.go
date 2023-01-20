package elements

// The SearchExpression element is an abstract element that represents the substituted element within a restriction. All search expressions derive from this base type. This element is not used in an XML instance document.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchexpression
import "encoding/xml"

type SearchExpression struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}

func (S *SearchExpression) SetForMarshal() {
	S.XMLName.Local = "t:SearchExpression"
}

func (S *SearchExpression) GetSchema() *Schema {
	return &SchemaTypes
}

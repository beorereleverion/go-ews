package goews

import (
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
)

// The CreateItem operation creates items in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitem-operation
func (c *client) CreateItem(eItem *elements.CreateItem) (*elements.CreateItemResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.CreateItemResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}

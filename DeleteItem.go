package goews

import (
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
)

// The DeleteItem operation deletes items in the Exchange store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitem-operation
func (c *client) DeleteItem(eItem *elements.DeleteItem) (*elements.DeleteItemResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.DeleteItemResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}

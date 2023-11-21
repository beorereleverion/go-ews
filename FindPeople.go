package goews

import (
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
)

// Find information about the FindPeople EWS operation.
// The FindPeople operation returns all persona objects from a specified Contacts folder or retrieves contacts that match a specified query string.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findpeople-operation
func (c *client) FindPeople(eItem *elements.FindPeople) (*elements.FindPeopleResponse, error) {
	oem, err := operations.NewEnvelopeMarshal(eItem)
	if err != nil {
		return nil, err
	}
	eResponse := &elements.FindPeopleResponse{}
	err = c.DoRequest(oem, eResponse)
	if err != nil {
		return nil, err
	}
	return eResponse, nil
}

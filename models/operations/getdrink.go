// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	"net/http"
)

type GetDrinkRequest struct {
	Name string `pathParam:"style=simple,explode=false,name=name"`
}

func (o *GetDrinkRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type GetDrinkResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A drink.
	Drink *components.Drink
	// An unknown error occurred interacting with the API.
	Error *components.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetDrinkResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDrinkResponse) GetDrink() *components.Drink {
	if o == nil {
		return nil
	}
	return o.Drink
}

func (o *GetDrinkResponse) GetError() *components.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetDrinkResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDrinkResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

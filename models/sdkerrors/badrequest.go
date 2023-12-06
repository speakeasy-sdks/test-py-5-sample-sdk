// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type BadRequest struct {
	// Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Error_ *string `json:"error,omitempty"`
	// HTTP status code
	StatusCode *float64 `json:"status_code,omitempty"`
	// The type of error returned
	TypeName *string `json:"type_name,omitempty"`
}

var _ error = &BadRequest{}

func (e *BadRequest) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

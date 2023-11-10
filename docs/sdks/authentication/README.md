# Authentication
(*Authentication*)

## Overview

The authentication endpoints.

### Available Operations

* [Authenticate](#authenticate) - Authenticate with the API by providing a username and password.

## Authenticate

Authenticate with the API by providing a username and password.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
)

func main() {
    s := testryan3.New(
        testryan3.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Authentication.Authenticate(ctx, operations.AuthenticateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AuthenticateRequestBody](../../models/operations/authenticaterequestbody.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.AuthenticateResponse](../../models/operations/authenticateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

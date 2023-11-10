# Authentication
(*Authentication*)

## Overview

The authentication endpoints.

### Available Operations

* [Login](#login) - Authenticate with the API by providing a username and password.

## Login

Authenticate with the API by providing a username and password.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
)

func main() {
    s := testryan3.New()


    operationSecurity := operations.LoginSecurity{
            Password: "<PASSWORD>",
            Username: "<USERNAME>",
        }

    ctx := context.Background()
    res, err := s.Authentication.Login(ctx, operations.LoginRequestBody{
        Type: operations.TypeAPIKey,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.LoginRequestBody](../../models/operations/loginrequestbody.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.LoginSecurity](../../models/operations/loginsecurity.md)       | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.LoginResponse](../../models/operations/loginresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

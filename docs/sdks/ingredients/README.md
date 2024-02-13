# Ingredients
(*Ingredients*)

## Overview

The ingredients endpoints.

### Available Operations

* [ListIngredients](#listingredients) - Get a list of ingredients.

## ListIngredients

Get a list of ingredients, if authenticated this will include stock levels and product codes otherwise it will only include public information.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"context"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
	"log"
)

func main() {
    s := testryan3.New(
        testryan3.WithSecurity(components.Security{
            APIKey: testryan3.String("<YOUR_API_KEY>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Ingredients.ListIngredients(ctx, operations.ListIngredientsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListIngredientsRequest](../../models/operations/listingredientsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListIngredientsResponse](../../models/operations/listingredientsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

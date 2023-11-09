# Orders
(*Orders*)

## Overview

The orders endpoints.

### Available Operations

* [CreateOrder](#createorder) - Create an order.

## CreateOrder

Create an order for a drink.

### Example Usage

```go
package main

import(
	"context"
	"log"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
	"github.com/speakeasy-sdks/test-ryan-3/models/callbacks"
)

func main() {
    s := testryan3.New(
        testryan3.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
        RequestBody: []components.OrderInput{
            components.OrderInput{
                ProductCode: "APM-1F2D3",
                Quantity: 26535,
                Type: components.OrderTypeDrink,
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateOrderRequest](../../models/operations/createorderrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateOrderResponse](../../models/operations/createorderresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |

<!-- Start SDK Example Usage [usage] -->
### Sign in

First you need to send an authentication request to the API by providing your username and password.
In the request body, you should specify the type of token you would like to receive: API key or JSON Web Token.
If your credentials are valid, you will receive a token in the response object: `res.object.token: str`.

```go
package main

import (
	"context"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
	"log"
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

### Browse available drinks

Once you are authenticated, you can use the token you received for all other authenticated endpoints.
For example, you can filter the list of available drinks by type.

```go
package main

import (
	"context"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
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
	res, err := s.Drinks.ListDrinks(ctx, operations.ListDrinksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```

### Create an order

When you submit an order, you can include a callback URL along your request.
This URL will get called whenever the supplier updates the status of your order.

```go
package main

import (
	"context"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
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
	res, err := s.Orders.CreateOrder(ctx, operations.CreateOrderRequest{
		RequestBody: []components.OrderInput{
			components.OrderInput{
				ProductCode: "APM-1F2D3",
				Quantity:    26535,
				Type:        components.OrderTypeDrink,
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

### Subscribe to webhooks to receive stock updates

```go
package main

import (
	"context"
	testryan3 "github.com/speakeasy-sdks/test-ryan-3"
	"github.com/speakeasy-sdks/test-ryan-3/models/components"
	"github.com/speakeasy-sdks/test-ryan-3/models/operations"
	"log"
	"net/http"
)

func main() {
	s := testryan3.New(
		testryan3.WithSecurity(components.Security{
			APIKey: testryan3.String("<YOUR_API_KEY>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Config.SubscribeToWebhooks(ctx, []operations.RequestBody{
		operations.RequestBody{},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->
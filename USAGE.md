<!-- Start SDK Example Usage -->
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
		testryan3.WithSecurity(""),
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
<!-- End SDK Example Usage -->
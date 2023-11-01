<!-- Start SDK Example Usage -->


```typescript
import { TheSpeakeasyBar } from "The-Speakeasy-Bar";
import { DrinkType } from "The-Speakeasy-Bar/dist/sdk/models/shared";

(async () => {
    const sdk = new TheSpeakeasyBar({
        apiKey: "",
    });

    const res = await sdk.drinks.listDrinks({});

    if (res.statusCode == 200) {
        // handle response
    }
})();

```
<!-- End SDK Example Usage -->